package contractor

import (
	ContractorAutoGen "github.com/t3kton/contractor_client/go/autogen"
)

// Contractor struct
type Contractor struct {
	*ContractorAutoGen.Contractor
}

// NewContractor creates a new Contractor and logs it in with the username and password
func NewContractor(host string, proxy string, username string, password string) (*Contractor, error) {
	c, err := ContractorAutoGen.NewContractor(host, proxy)
	if err != nil {
		return nil, err
	}

	c2 := Contractor{}
	c2.Contractor = c

	token, err := c.AuthUserCallLogin(username, password)
	if err != nil {
		return nil, err
	}

	c2.SetHeader("Auth-Id", username)
	c2.SetHeader("Auth-Token", token)

	return &c2, nil
}

// Logout calles logout and removes the auth headers
func (c *Contractor) Logout() {
	c.AuthUserCallLogout()

	c.ClearHeader("Auth-Id")
	c.ClearHeader("Auth-Token")
}
