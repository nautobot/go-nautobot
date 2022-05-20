package nautobot_test

import (
	"context"
	"testing"

	nb "github.com/nautobot/go-nautobot/pkg/nautobot"
	"github.com/tidwall/gjson"
)

const (
	deviceExist = "manufacturer with this name already exists."
)

func TestDcimManufacturersCreateWithResponse(t *testing.T) {
	nb = get_nb_session(t)

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

			var m nb.Manufacturer
			m.Name = tc.vendor

			rsp, err := c.DcimManufacturersCreateWithResponse(
				context.Background(),
				nb.DcimManufacturersCreateJSONRequestBody(m))

			data := string(rsp.Body)
			dataName := gjson.Get(data, "name.0")

			if dataName.String() == deviceExist && tc.err == deviceExist {
				return
			}

			if err != nil {
				t.Fatalf("failed to create manufacturer %s on %s: %s", tc.vendor, server, err.Error())
			}
		})
	}

}
