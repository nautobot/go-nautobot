package main

import (
	"context"
	"fmt"

	nb "github.com/nautobot/go-nautobot/pkg/nautobot"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	token, err := nb.NewSecurityProviderNautobotToken("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	check(err)

	c, err := nb.NewClientWithResponses(
		"https://demo.nautobot.com/api/",
		nb.WithRequestEditorFn(token.Intercept),
	)
	check(err)

	ctx := context.Background()

	resp, err := c.DcimManufacturersListWithResponse(ctx, &nb.DcimManufacturersListParams{})
	check(err)

	fmt.Printf("%v", string(resp.Body))

}
