package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	ttemplate "text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// TerraformConfig holds the Terraform configuration parameters
type TerraformConfig struct {
	Token      string
	K8sVersion string
	Label      string
	Region     string
	Tags       []string
	Pools      []map[string]interface{}
}

// ProvisionRequest is used to decode JSON request for provisioning
type ProvisionRequest struct {
	Platform string `json:"platform"`
	APIKey   string `json:"apiKey"`
}

// Constants and variables
const terraformTemplate = `token = "{{ .Token }}"`

func generateTerraformFile(config TerraformConfig) (string, error) {
	tmpl, err := ttemplate.New("terraform").Parse(terraformTemplate)
	if err != nil {
		return "", fmt.Errorf("error parsing Terraform template: %w", err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, config)
	if err != nil {
		return "", fmt.Errorf("error executing Terraform template: %w", err)
	}

	return buffer.String(), nil
}

func isBase64(s string) bool {
	// Check if the string length is a multiple of 4
	// This check is relaxed to allow for optional padding.
	if len(s)%4 != 0 && (len(s)+1)%4 != 0 && (len(s)+2)%4 != 0 {
		return false
	}

	// Check if the string contains only valid Base64 characters
	for _, r := range s {
		if !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '+' || r == '/' || r == '=') {
			return false
		}
	}

	// Optionally, try decoding it to check if it's valid Base64
	// Note: This will also return true for strings that are technically valid Base64 but weren't necessarily encoded from binary data
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Template setup
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("error parsing templates:", err)
	}

	// Define routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Your homepage handler logic here
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		// Your login handler logic here
		tmpl.ExecuteTemplate(w, "login.html", nil)
	})

	r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
		// redirect to provider.html
		http.Redirect(w, r, "/create", http.StatusSeeOther)
	})

	r.Get("/create", func(w http.ResponseWriter, r *http.Request) {
		// Your provider handler logic here
		tmpl.ExecuteTemplate(w, "create.html", nil)
	})

	r.Post("/provision", provisionHandler)

	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("templates/assets"))))

	// Start the server
	fmt.Println("Starting server on http://127.0.0.1:8080")
	http.ListenAndServe(":8080", r)
}

func provisionHandler(w http.ResponseWriter, r *http.Request) {
	var request ProvisionRequest
	// Decode JSON request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// // Validate API key
	// if request.APIKey != "your_valid_api_key" {
	// 	http.Error(w, "Invalid API key", http.StatusUnauthorized)
	// 	return
	// }

	// Assume TerraformConfig is somehow obtained or constructed here
	tfConfig := TerraformConfig{
		Token: request.APIKey,
	}

	// Generate Terraform file
	tfFileContent, err := generateTerraformFile(tfConfig)
	if err != nil {
		http.Error(w, "Error generating Terraform file", http.StatusInternalServerError)
		return
	}

	// Create Terraform directory if it doesn't exist
	if _, err := os.Stat("terraform"); os.IsNotExist(err) {
		if err := os.Mkdir("terraform", 0755); err != nil {
			http.Error(w, "Error creating Terraform directory", http.StatusInternalServerError)
			return
		}
	}

	// cleanup all files in terraform directory
	files, err := os.ReadDir("terraform")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if err := os.RemoveAll("terraform/" + file.Name()); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Platform: ", request.Platform)

	tempFiles, err := os.ReadDir(request.Platform)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range tempFiles {
		// copy all .tf files to terraform directory
		if strings.HasSuffix(file.Name(), ".tf") {
			data, err := os.ReadFile(request.Platform + "/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			if err := os.WriteFile("terraform/"+file.Name(), data, 0644); err != nil {
				log.Fatal(err)
			}
		}
	}

	// Save .tf file
	if err := os.WriteFile("terraform/terraform.tfvars", []byte(tfFileContent), 0644); err != nil {
		http.Error(w, "Error saving Terraform file", http.StatusInternalServerError)
		return
	}

	// fmt.Fprintf(w, "Provisioning initiated")

	// Execute Terraform commands
	executeTerraformCommand(w, "terraform", "init")

	executeTerraformCommand(w, "terraform", "apply", "-auto-approve")

	// if no error from above, clear http.ResponseWriter and write kubeconfig to w

	// terraform output kubeconfig > kubeconfig.yaml

	executeTerraformCommand(w, "terraform", "output", "kubeconfig")

	// if w contains no error, write
}

func executeTerraformCommand(w http.ResponseWriter, command string, args ...string) {
	// initText := "Terraform has been successfully initialized!\n"
	cmd := exec.Command(command, args...)
	cmd.Dir = "terraform"

	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("Error executing %s: %s", command, err)
		log.Printf("Output: %s", out)
		http.Error(w, "Error executing Terraform command", http.StatusInternalServerError)

	} else {
		log.Printf("Successfully executed %s", command)
		log.Printf("Output: %s", out)
		// check if out is base64 format
		if isBase64(string(out)) {
			// remove first and last character from out
			if out[0] == '"' && out[len(out)-1] == '"' {
				out = out[1 : len(out)-1]
			}
			// decode base64 kubeconfig
			kubeconfig, err := base64.StdEncoding.DecodeString(string(out))
			if err != nil {
				log.Printf("Error decoding kubeconfig: %s", err)
			}
			w.Write(kubeconfig)
		} else {
			w.Write(out)
		}
	}
}
