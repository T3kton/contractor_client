// Package contractor - Automatically generated by cinp-codegen from http://127.0.0.1:8889/api/v1/ at 2022-01-08T19:31:37.129298
package contractor

import (
	"fmt"

	cinp "github.com/cinp/go"
)

// Contractor from http://127.0.0.1:8889/api/v1/
type Contractor struct {
	cinp *cinp.CInP
}

// NewContractor creates and returns a new Contractor
func NewContractor(host string, proxy string, username string, password string) (*Contractor, error) {
	var err error
	s := Contractor{}
	s.cinp, err = cinp.NewCInP(host, "/api/v1/", proxy)
	if err != nil {
		return nil, err
	}

	registerAuth(s.cinp)
	registerBluePrint(s.cinp)
	registerSite(s.cinp)
	registerSurvey(s.cinp)
	registerDirectory(s.cinp)
	registerUtilities(s.cinp)
	registerBuilding(s.cinp)
	registerForeman(s.cinp)
	registerSubContractor(s.cinp)
	registerPostOffice(s.cinp)
	registerRecords(s.cinp)
	registerVirtualBox(s.cinp)
	registerVCenter(s.cinp)
	registerDocker(s.cinp)
	registerIPMI(s.cinp)
	registerManual(s.cinp)
	registerProxmox(s.cinp)
	registerAzure(s.cinp)
	registerAMT(s.cinp)
	registerRaspberryPi(s.cinp)
	registerPacket(s.cinp)
	registerAWS(s.cinp)
	registerRedFish(s.cinp)

	APIVersion, err := s.GetAPIVersion("/api/v1/")
	if err != nil {
		return nil, err
	}

	if APIVersion != "0.9" {
		return nil, fmt.Errorf("API version mismatch.  Got '%s', expected '0.9'", APIVersion)
	}

	if username != "" {
		token, err := s.AuthUserCallLogin(username, password)
		if err != nil {
			return nil, err
		}
		s.cinp.SetAuth(username, token)
	}

	return &s, nil
}

// Logout calles the Logout method and clears the auth token
func (s *Contractor) Logout() {
	if s.cinp.IsAuthenticated() {
		s.AuthUserCallLogout()
		s.cinp.SetAuth("", "")
	}
}

// GetAPIVersion Get the API version number for the Namespace at the URI
func (s *Contractor) GetAPIVersion(uri string) (string, error ) {
	r, t, err := s.cinp.Describe(uri)
	if err != nil {
		return "", err
	}

	if t != "Namespace" {
		return "", fmt.Errorf("Excpected a Namespace got '%s'", t)
	}

	return r.APIVersion, nil
}