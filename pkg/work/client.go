package client

import (
	"context"
	"log"
	"os"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
)

func Run(ctx context.Context) error {

	return nil
}

func newManager(ctx context.Context) error {
	organizationUrl := os.Getenv("ADO_ORG")
	pat := os.Getenv("ADO_PAT")
	if organizationUrl == "" || pat == "" {
		log.Fatal("useless error message")
	}

	// Create a connection to your organization
	connection := azuredevops.NewPatConnection(organizationUrl, pat)

	wicli, err := workitemtracking.NewClient(ctx, connection)

	if err != nil {
		log.Fatal(err)
	}

	workItem := workitemtracking.CreateWorkItemArgs{}

	wicli.CreateWorkItem(ctx, workItem)

	return nil
}

/*

type CreateWorkItemArgs struct {
		// (required) The JSON Patch document representing the work item
		Document *[]webapi.JsonPatchOperation
		// (required) Project ID or project name
		Project *string
		// (required) The work item type of the work item to create
		Type *string
		// (optional) Indicate if you only want to validate the changes without saving the work item
		ValidateOnly *bool
		// (optional) Do not enforce the work item type rules on this update
		BypassRules *bool
		// (optional) Do not fire any notifications for this change
		SuppressNotifications *bool
		// (optional) The expand parameters for work item attributes. Possible options are { None, Relations, Fields, Links, All }.
		Expand *WorkItemExpand
	}
*/
