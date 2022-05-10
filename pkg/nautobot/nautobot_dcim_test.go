package nautobot

import (
	"context"
	"testing"

	"github.com/tidwall/gjson"
)

const (
	deviceExist = "manufacturer with this name already exists."
)

func TestDcimManufacturersCreateWithResponse(t *testing.T) {
	c, err := CreateNautobotSession()
	if err != nil {
		t.Fatal("Failed to create Nautobot Session", err.Error())
	}

	// Test table
	tt := []struct {
		name   string
		vendor string
		err    string
	}{
		{name: "New Vendor", vendor: "new vendor 1"},
		{name: "Existing Vendor", vendor: "Cisco", err: deviceExist},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			var m Manufacturer
			m.Name = tc.vendor

			rsp, err := c.DcimManufacturersCreateWithResponse(
				context.Background(),
				DcimManufacturersCreateJSONRequestBody(m))

			data := string(rsp.Body)
			dataName := gjson.Get(data, "name.0")

			if dataName.String() == deviceExist && tc.err == deviceExist {
				return
			}

			if err != nil {
				t.Fatalf("failed to create manufacturer %s: %s", tc.vendor, err.Error())
			}
		})
	}

}
