package nautobot

import (
	"context"
	"testing"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/tidwall/gjson"
)

func TestManufacturerBulkOperations(t *testing.T) {
	nautobotSession, err := CreateNautobotSession()
	if err != nil {
		t.Fatalf("Failed to create Nautobot Session: %s", err.Error())
	}

    descriptions := []string{"A test manufacturer.", "An example manufacturer."}
	var manufacturers = []BulkWritableManufacturerRequest {
	    BulkWritableManufacturerRequest {
	        Name: "Test Vendor Inc.",
	        Description: &descriptions[0],
	    },
	    BulkWritableManufacturerRequest {
	        Name: "Example Vendor Ltd.",
	        Description: &descriptions[1],
	    },
	}
	var updated_description = "An updated description"

    for index, manufacturer := range manufacturers {
        var manufacturer_object = ManufacturerRequest{Name: manufacturer.Name, Description: manufacturer.Description}
        creation_response, creation_error := nautobotSession.DcimManufacturersCreateWithResponse(
            context.Background(),
            DcimManufacturersCreateJSONRequestBody(manufacturer_object))
        if creation_error != nil {
            t.Fatalf("failed to create manufacturer: %s", creation_error.Error())
        }
        if creation_response.StatusCode() / 100 != 2 {
            t.Fatalf("Failed to create manufacturer, got status code %s and error %s", creation_response.Status(), creation_response.Body)
        }
        data := string(creation_response.Body)
        manufacturers[index].Id = openapi_types.UUID(gjson.Get(data, "id").String())
        manufacturers[index].Description = &updated_description
    }

	bulk_update_response, bulk_update_err := nautobotSession.DcimManufacturersBulkUpdateWithResponse(
	    context.Background(),
	    DcimManufacturersBulkUpdateJSONRequestBody(manufacturers),
	)

    if bulk_update_err != nil {
        t.Fatalf("Failed to bulk update manufacturers: %s", bulk_update_err.Error())
    }
	if bulk_update_response.StatusCode() / 100 != 2 {
		t.Fatalf("Failed to bulk update manufacturers, got status code %s and error %s", bulk_update_response.Status(), bulk_update_response.Body)
	}

    body_as_json := gjson.Parse(string(bulk_update_response.Body))
    index := 0
    var manufacturers_to_destroy [2]BulkOperationRequest
    body_as_json.ForEach(func(_, value gjson.Result) bool {
	    var expected_description = *manufacturers[index].Description
	    var actual_description = value.Get("description").String()
	    if expected_description != actual_description {
	        t.Fatalf("Failed to bulk update manufacturer description, expected %s - got %s", expected_description, actual_description)
	    }
	    manufacturers_to_destroy[index] = BulkOperationRequest{Id: manufacturers[index].Id}
	    index++
	    return true
    })

    bulk_delete_response, bulk_delete_err := nautobotSession.DcimManufacturersBulkDestroyWithResponse(
	    context.Background(),
	    DcimManufacturersBulkDestroyJSONRequestBody(manufacturers_to_destroy[:]),
	)

    if bulk_delete_err != nil {
        t.Fatalf("Failed to bulk delete manufacturers: %s", bulk_delete_err.Error())
    }
	if bulk_delete_response.StatusCode() / 100 != 2 {
		t.Fatalf("Failed to bulk delete manufacturers, got status code %s and error %s", bulk_delete_response.Status(), bulk_delete_response.Body)
	}
}

func TestManufacturerBasicOperations(t *testing.T) {
	nautobotSession, err := CreateNautobotSession()
	if err != nil {
		t.Fatalf("Failed to create Nautobot Session: %s", err.Error())
	}

	// Prepare Manufacturer record to create
	var manufacturer ManufacturerRequest
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
