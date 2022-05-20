package nautobot_test

import (
	"fmt"
	"os"
	"testing"

	nb "github.com/nautobot/go-nautobot/pkg/nautobot"
)

func getenv(name string) (string, error) {
	v := os.Getenv(name)
	if v == "" {
		return v, fmt.Errorf("%s environment variable not set", name)
	}
	return v, nil
}

func get_nb_session(t *testing.T) {
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
	return nb
}
