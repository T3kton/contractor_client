package contractor

import (
	"context"
	"log/slog"
)

// NewContractor creates a new Contractor and logs it in with the username and password
func NewContractor(ctx context.Context, log *slog.Logger, host string, proxy string, username string, password string) (*Contractor, error) {
	c, err := NewContractorInt(ctx, log, host, proxy)
	if err != nil {
		return nil, err
	}

	token, err := c.AuthUserCallLogin(ctx, username, password)
	if err != nil {
		return nil, err
	}

	c.SetHeader("Auth-Id", username)
	c.SetHeader("Auth-Token", token)

	return c, nil
}

// Logout calles logout and removes the auth headers
func (c *Contractor) Logout(ctx context.Context) {
	c.AuthUserCallLogout(ctx)

	c.ClearHeader("Auth-Id")
	c.ClearHeader("Auth-Token")
}
