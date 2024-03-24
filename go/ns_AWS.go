// Package contractor - (version: "0.1") - Automatically generated by cinp-codegen from /api/v1/AWS/ at 2024-03-24T00:03:41.412767
package contractor

import (
	cinp "github.com/cinp/go"
	"fmt"
	"reflect"
	"time"
)
// AwsAWSEC2Foundation - Model AWSEC2Foundation(/api/v1/AWS/AWSEC2Foundation)
/*
AWSEC2Foundation(locator, site, blueprint, id_map, located_at, built_at, updated, created, foundation_ptr, awsec2_instance_id)
*/
type AwsAWSEC2Foundation struct {
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
	Awsec2InstanceID string `json:"awsec2_instance_id"`
	State string `json:"state"`
	Type string `json:"type"`
	ClassList string `json:"class_list"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *AwsAWSEC2Foundation) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"locator": object.Locator,
			"site": object.Site,
			"blueprint": object.Blueprint,
			"id_map": object.IDMap,
			"awsec2_instance_id": object.Awsec2InstanceID,
		}
	}
	return &map[string]interface{}{ 
		"site": object.Site,
		"blueprint": object.Blueprint,
		"id_map": object.IDMap,
		"awsec2_instance_id": object.Awsec2InstanceID,
	}
}

// AwsAWSEC2FoundationNew - Make a new object of Model AWSEC2Foundation
func (service *Contractor) AwsAWSEC2FoundationNew() *AwsAWSEC2Foundation {
	return &AwsAWSEC2Foundation{cinp: service.cinp}
}

// AwsAWSEC2FoundationNewWithID - Make a new object of Model AWSEC2Foundation
func (service *Contractor) AwsAWSEC2FoundationNewWithID(id string) *AwsAWSEC2Foundation {
	result := AwsAWSEC2Foundation{cinp: service.cinp}
	result.SetID("/api/v1/AWS/AWSEC2Foundation:" + id + ":")
	return &result
}

// AwsAWSEC2FoundationGet - Get function for Model AWSEC2Foundation
func (service *Contractor) AwsAWSEC2FoundationGet(id string) (*AwsAWSEC2Foundation, error) {
	object, err := service.cinp.Get("/api/v1/AWS/AWSEC2Foundation:" + id + ":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*AwsAWSEC2Foundation)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model AWSEC2Foundation
func (object *AwsAWSEC2Foundation) Create() error {
	if err := object.cinp.Create("/api/v1/AWS/AWSEC2Foundation", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model AWSEC2Foundation
func (object *AwsAWSEC2Foundation) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model AWSEC2Foundation
func (object *AwsAWSEC2Foundation) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// AwsAWSEC2FoundationListFilters - Return a slice of valid filter names AWSEC2Foundation
func (service *Contractor) AwsAWSEC2FoundationListFilters() [1]string {
  return [1]string{ "site" }
}

// AwsAWSEC2FoundationList - List function for Model AWSEC2Foundation
func (service *Contractor) AwsAWSEC2FoundationList(filterName string, filterValues map[string]interface{}) (<-chan *AwsAWSEC2Foundation, error) {
	if filterName != "" {
		for _, item := range service.AwsAWSEC2FoundationListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("Filter '%s' is invalid", filterName)
	}
	good:

	in := service.cinp.ListObjects("/api/v1/AWS/AWSEC2Foundation", reflect.TypeOf(AwsAWSEC2Foundation{}), filterName, filterValues, 50)
	out := make(chan *AwsAWSEC2Foundation)
	go func() {
		defer close(out)
		for v := range in {
			v.(*AwsAWSEC2Foundation).cinp = service.cinp
			out <- v.(*AwsAWSEC2Foundation)
		}
	}()
	return out, nil
}

func registerAWS(cinp *cinp.CInP) {
	cinp.RegisterType("/api/v1/AWS/AWSEC2Foundation", reflect.TypeOf((*AwsAWSEC2Foundation)(nil)).Elem())
}
