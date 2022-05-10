package nautobot

import (
	"context"
	"fmt"
	"net/http"
)

type SecurityProviderNautobotToken struct {
	token string
}

func NewSecurityProviderNautobotToken(t string) (*SecurityProviderNautobotToken, error) {
	return &SecurityProviderNautobotToken{
		token: t,
	}, nil
}

func (s *SecurityProviderNautobotToken) Intercept(ctx context.Context, req *http.Request) error {
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", s.token))
	return nil
}
