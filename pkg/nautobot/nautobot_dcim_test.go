package nautobot

import (
	"context"
	"testing"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/tidwall/gjson"
)

func TestManufacturerBasicOperations(t *testing.T) {
	nautobotSession, err := CreateNautobotSession()
	if err != nil {
		t.Fatalf("Failed to create Nautobot Session: %s", err.Error())
	}

	// Prepare Manufacturer record to create
	var manufacturer Manufacturer
	manufacturer.Name = "Nautobot Technologies Inc."

	// Create Manufacturer record
	creation_response, creation_err := nautobotSession.DcimManufacturersCreateWithResponse(
		context.Background(),
		DcimManufacturersCreateJSONRequestBody(manufacturer))

	// Check creation for errors
	if creation_err != nil {
		t.Fatalf("Failed to create a manufacturer %s: %s", manufacturer.Name, creation_err.Error())
	}
	if creation_response.StatusCode() / 100 != 2 {
		t.Fatalf("Failed to create a manufacturer, got status code %s and error %s", creation_response.Status(), creation_response.Body)
	}

	// Extract the UUID of the created manufacturer for later use
	responseData := string(creation_response.Body)
	dataUUID := openapi_types.UUID(gjson.Get(responseData, "id").String())

	// Prepare Manufacturer record for updating
	updated_description := "This is a testing manufacturer."
	manufacturer.Description = &updated_description

	// Update Manufacturer record
    update_response, update_err := nautobotSession.DcimManufacturersUpdateWithResponse(
        context.Background(),
        dataUUID,
        DcimManufacturersUpdateJSONRequestBody(manufacturer))

    // Check update for errors
	if update_err != nil {
		t.Fatalf("Failed to update a manufacturer %s: %s", manufacturer.Name, update_err.Error())
	}
	if update_response.StatusCode() / 100 != 2 {
		t.Fatalf("Failed to update a manufacturer, got status code %s and error %s", update_response.Status(), update_response.Body)
	}

	// Delete Manufacturer record
    delete_response, delete_err := nautobotSession.DcimManufacturersDestroyWithResponse(
        context.Background(),
        dataUUID)

    // Check deletion for errors
	if delete_err != nil {
		t.Fatalf("Failed to delete a manufacturer %s: %s", manufacturer.Name, delete_err.Error())
	}
	if delete_response.StatusCode() / 100 != 2 {
		t.Fatalf("Failed to delete a manufacturer, got status code %s and error %s", delete_response.Status(), delete_response.Body)
	}
}
