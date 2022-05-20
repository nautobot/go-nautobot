package nautobot

import (
	"fmt"
	"os"
)

func getenv(name string) (string, error) {
	v := os.Getenv(name)
	if v == "" {
		return v, fmt.Errorf("%s environment variable not set", name)
	}
	return v, nil
}

func CreateNautobotSession() (c *ClientWithResponses, e error) {
	server, err := getenv("NAUTOBOT_URL")
	if err != nil {
		return nil, err
	}
	token, err := getenv("NAUTOBOT_TOKEN")
	if err != nil {
		return nil, err
	}

	tk, err := NewSecurityProviderNautobotToken(token)
	if err != nil {
		return nil, err
	}

	c, err = NewClientWithResponses(
		server,
		WithRequestEditorFn(tk.Intercept),
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}
