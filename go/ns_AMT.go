// Package contractor - (version: "0.1") - Automatically generated by cinp-codegen from /api/v1/AMT/ at 2024-03-24T00:03:41.412767
package contractor

import (
	cinp "github.com/cinp/go"
	"fmt"
	"reflect"
	"time"
)
// AmtAMTFoundation - Model AMTFoundation(/api/v1/AMT/AMTFoundation)
/*
AMTFoundation(locator, site, blueprint, id_map, located_at, built_at, updated, created, foundation_ptr, amt_username, amt_password, amt_ip_address, plot)
*/
type AmtAMTFoundation struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Locator string `json:"locator"`
	Site string `json:"site"`
	Blueprint string `json:"blueprint"`
	IDMap string `json:"id_map"`
	LocatedAt time.Time `json:"located_at"`
	BuiltAt time.Time `json:"built_at"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
	AmtUsername string `json:"amt_username"`
	AmtPassword string `json:"amt_password"`
	AmtIPAddress string `json:"amt_ip_address"`
	Plot string `json:"plot"`
	State string `json:"state"`
	Type string `json:"type"`
	ClassList string `json:"class_list"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *AmtAMTFoundation) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"locator": object.Locator,
			"site": object.Site,
			"blueprint": object.Blueprint,
			"id_map": object.IDMap,
			"amt_username": object.AmtUsername,
			"amt_password": object.AmtPassword,
			"amt_ip_address": object.AmtIPAddress,
			"plot": object.Plot,
		}
	}
	return &map[string]interface{}{ 
		"site": object.Site,
		"blueprint": object.Blueprint,
		"id_map": object.IDMap,
		"amt_username": object.AmtUsername,
		"amt_password": object.AmtPassword,
		"amt_ip_address": object.AmtIPAddress,
		"plot": object.Plot,
	}
}

// AmtAMTFoundationNew - Make a new object of Model AMTFoundation
func (service *Contractor) AmtAMTFoundationNew() *AmtAMTFoundation {
	return &AmtAMTFoundation{cinp: service.cinp}
}

// AmtAMTFoundationNewWithID - Make a new object of Model AMTFoundation
func (service *Contractor) AmtAMTFoundationNewWithID(id string) *AmtAMTFoundation {
	result := AmtAMTFoundation{cinp: service.cinp}
	result.SetID("/api/v1/AMT/AMTFoundation:" + id + ":")
	return &result
}

// AmtAMTFoundationGet - Get function for Model AMTFoundation
func (service *Contractor) AmtAMTFoundationGet(id string) (*AmtAMTFoundation, error) {
	object, err := service.cinp.Get("/api/v1/AMT/AMTFoundation:" + id + ":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*AmtAMTFoundation)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model AMTFoundation
func (object *AmtAMTFoundation) Create() error {
	if err := object.cinp.Create("/api/v1/AMT/AMTFoundation", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model AMTFoundation
func (object *AmtAMTFoundation) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model AMTFoundation
func (object *AmtAMTFoundation) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// AmtAMTFoundationListFilters - Return a slice of valid filter names AMTFoundation
func (service *Contractor) AmtAMTFoundationListFilters() [1]string {
  return [1]string{ "site" }
}

// AmtAMTFoundationList - List function for Model AMTFoundation
func (service *Contractor) AmtAMTFoundationList(filterName string, filterValues map[string]interface{}) (<-chan *AmtAMTFoundation, error) {
	if filterName != "" {
		for _, item := range service.AmtAMTFoundationListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("Filter '%s' is invalid", filterName)
	}
	good:

	in := service.cinp.ListObjects("/api/v1/AMT/AMTFoundation", reflect.TypeOf(AmtAMTFoundation{}), filterName, filterValues, 50)
	out := make(chan *AmtAMTFoundation)
	go func() {
		defer close(out)
		for v := range in {
			v.(*AmtAMTFoundation).cinp = service.cinp
			out <- v.(*AmtAMTFoundation)
		}
	}()
	return out, nil
}

func registerAMT(cinp *cinp.CInP) {
	cinp.RegisterType("/api/v1/AMT/AMTFoundation", reflect.TypeOf((*AmtAMTFoundation)(nil)).Elem())
}
