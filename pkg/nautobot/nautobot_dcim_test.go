package nautobot

import (
	"context"
	"testing"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/tidwall/gjson"
)

const (
	deviceExist    = "manufacturer with this name already exists."
	deviceNotExist = "manufacturer with this id doesn't exist."
)

func TestDcimManufacturersDestroyWithResponse(t *testing.T) {
	nautobotSession, err := CreateNautobotSession()
	if err != nil {
		t.Fatalf("Failed to create Nautobot Session: %s", err.Error())
	}

	// Prepare Manufacturer to delete
	var manufacturer Manufacturer
	manufacturer.Name = "Existing-Manufacturer"
	response, err := nautobotSession.DcimManufacturersCreateWithResponse(
		context.Background(),
		DcimManufacturersCreateJSONRequestBody(manufacturer))
	if err != nil {
		t.Fatalf("Failed to create a manufacturer %s: %s", manufacturer.Name, err.Error())
	}
	if response.StatusCode() != 201 {
		t.Fatalf("Failed to create a manufacturer, got status code %s", response.Status())
	}
	responseData := string(response.Body)

	// Extract the UUID of the created manufacturer for later use
	dataUUID := openapi_types.UUID(gjson.Get(responseData, "id").String())

	testArray := []struct {
		name       string
		uuid       openapi_types.UUID
		statusCode int
	}{
		{name: "Delete an existing Manufacturer", uuid: dataUUID, statusCode: 204},
		{name: "Try to delete a non-existing Manufacturer", uuid: "1234", statusCode: 404},
	}

	for _, testCase := range testArray {
		t.Run(testCase.name, func(t *testing.T) {
			response, err := nautobotSession.DcimManufacturersDestroyWithResponse(
				context.Background(),
				testCase.uuid)

			if err != nil {
				t.Fatalf("Failed to delete manufacturer %s: %s", testCase.name, err.Error())
			}

			if response.StatusCode() == testCase.statusCode {
				return
			} else {
				responseData := string(response.Body)
				t.Fatalf("Expected %d, got %d: Failed to delete manufacturer '%s': %s", testCase.statusCode, response.StatusCode(), testCase.name, responseData)
			}

		})
	}
}

func TestDcimManufacturersCreateWithResponse(t *testing.T) {
	nautobotSession, err := CreateNautobotSession()
	if err != nil {
		t.Fatal("Failed to create Nautobot Session", err.Error())
	}

	testArray := []struct {
		name   string
		vendor string
		err    string
	}{
		{name: "New Vendor", vendor: "new vendor 1"},
		{name: "Existing Vendor", vendor: "Cisco", err: deviceExist},
	}

	for _, testCase := range testArray {
		t.Run(testCase.name, func(t *testing.T) {

			var manufacturer Manufacturer
			manufacturer.Name = testCase.vendor

			response, err := nautobotSession.DcimManufacturersCreateWithResponse(
				context.Background(),
				DcimManufacturersCreateJSONRequestBody(manufacturer))

			responseData := string(response.Body)
			dataName := gjson.Get(responseData, "name.0")

			if dataName.String() == deviceExist && testCase.err == deviceExist {
				return
			}

			if response.StatusCode() != 201 {
				t.Fatalf("Failed to create a manufacturer %s: %s", testCase.vendor, response.Status())
			}

			if err != nil {
				t.Fatalf("Failed to create manufacturer %s: %s", testCase.vendor, err.Error())
			}
		})
	}

}
