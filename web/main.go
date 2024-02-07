package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

// define port number
const host = "localhost"
const port = ":8080"
const conf = "../conf.d"

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/run", runTerraformHandler)

	fmt.Printf("Terraform webserver started at http://%s%s\n", host, port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Terraform Runner!")
}

func runTerraformHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize Terraform
	cmd := exec.Command("terraform", "init")
	cmd.Dir = conf
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to initialize Terraform: %v", err)
	}

	// Apply Terraform
	cmd = exec.Command("terraform", "apply", "-auto-approve")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to apply Terraform: %v", err)
	}

	fmt.Fprintf(w, "Terraform script executed successfully.")
}
