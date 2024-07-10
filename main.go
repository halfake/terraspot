package main

import (
	"fmt"
	"log"
	"os"

	vault "github.com/hashicorp/vault/api"
)

var GitLab_resources = [...]string{"projects", "users"}
var Vault_resources = [...]string{"policies", "roles"}

func gitlab_instance() {

}

func vault_instance() {
	config := vault.DefaultConfig()

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// Authenticate with a token
	client.SetToken(os.Getenv("VAULT_TOKEN"))

	authBackends, err := client.Sys().ListAuth()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Roles retrieved:", authBackends)

	for path, auth := range authBackends {
		fmt.Printf("Auth Backend Path: %s, Type: %s\n", path, auth.Type)

		// Construct the path to list roles
		rolesPath := fmt.Sprintf("%s/roles", path)

		// Read roles
		roles, err := client.Logical().List(rolesPath)
		if err != nil {
			log.Printf("Unable to list roles for backend at %s: %v", path, err)
			continue
		}

		if roles == nil || roles.Data == nil {
			fmt.Printf("No roles found for backend at %s\n", path)
			continue
		}

		// Print the roles
		if keys, ok := roles.Data["keys"].([]interface{}); ok {
			fmt.Println("Roles:")
			for _, role := range keys {
				fmt.Printf(" - %s\n", role.(string))
			}
		} else {
			fmt.Printf("No roles found for backend at %s\n", path)
		}
	}
}

func terraform_state() {

}

func main() {

	// TODO Loop over Terraform projects and collect state
	terraform_state()

	// TODO Loop over instance, list resources via their API, and match against the Terraform global state compiled before
	gitlab_instance()
	vault_instance()
}
