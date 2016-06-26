package main

import (
	"fmt"

	"github.com/JesusTinoco/go-smith/stacksmith"
)

func main() {
	APIKey := "<API_KEY_STACKSMITH>"
	client := stacksmith.NewClient(APIKey, nil)

	pag := &stacksmith.PaginationParams{Page: 1, PerPage: 100}

	// Get all the stacks created
	stacksList, _, _ := client.Stacks.List(pag)
	fmt.Println(fmt.Sprintf("You have %d stacks.", len(stacksList.Items)))

	// Remove an stack
	stackToRemove := "<STACK_ID>"
	status, _, _ := client.Stacks.Delete(stackToRemove)
	if status.Deleted {
		fmt.Println(fmt.Sprintf("The stack with id %s has been deleted", stackToRemove))
	}

	// Get the vulnerabilities fron a given stack
	stackID := "<STACK_ID"
	vulnerabilities, _, _ := client.Stacks.GetVulnerabilities(stackID, pag)
	fmt.Println(fmt.Sprintf("The '%s' stack has %d vulnerabilities",
		stackID, len(vulnerabilities.Items)))
}
