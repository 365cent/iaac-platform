package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
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
// token = "0a5c8ce5613ae1be44952c0935f8ad19198d05c3f5d8564178bc1c4a634987f4"
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
		// copy all files from platform directory to terraform directory
		if err := os.WriteFile("terraform/"+file.Name(), []byte(file.Name()), 0644); err != nil {
			http.Error(w, "Error saving Terraform file", http.StatusInternalServerError)
			return
		}
	}

	// Save .tf file
	if err := os.WriteFile("terraform/terraform.tfvars", []byte(tfFileContent), 0644); err != nil {
		http.Error(w, "Error saving Terraform file", http.StatusInternalServerError)
		return
	}

	// Execute Terraform commands
	executeTerraformCommand("terraform", "init")
	executeTerraformCommand("terraform", "apply", "-auto-approve")

	fmt.Fprintf(w, "Provisioning initiated")
}

func executeTerraformCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Dir = "terraform"
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("Error executing %s: %s", command, err)
		log.Printf("Output: %s", out)
	} else {
		log.Printf("Successfully executed %s", command)
		log.Printf("Output: %s", out)
	}
}
