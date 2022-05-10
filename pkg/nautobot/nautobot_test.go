package nautobot_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	nb "github.com/nautobot/go-nautobot/pkg/nautobot"
	"github.com/tidwall/gjson"
)

const (
	deviceExist = "manufacturer with this name already exists."
)

func getenv(name string) (string, error) {
	v := os.Getenv(name)
	if v == "" {
		return v, fmt.Errorf("%s environment variable not set", name)
	}
	return v, nil
}

func TestDcimManufacturersCreateWithResponse(t *testing.T) {
	server, err := getenv("NAUTOBOT_URL")
	if err != nil {
		t.Fatal(err)
	}
	token, err := getenv("NAUTOBOT_TOKEN")
	if err != nil {
		t.Fatal(err)
	}

	tk, err := nb.NewSecurityProviderNautobotToken(token)
	if err != nil {
		t.Fatalf("could not create a Nautobot token for %s", server)
	}

	c, err := nb.NewClientWithResponses(
		server,
		nb.WithRequestEditorFn(tk.Intercept),
	)
	if err != nil {
		t.Fatalf("could not setup a client connection to  %s", server)
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
