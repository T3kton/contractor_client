// Package contractor - (version: "0.1") - Automatically generated by cinp-codegen from /api/v1/AWS/ at 2024-05-23T18:02:09.902032
package contractor

import (
	"context"
	"fmt"
	cinp "github.com/cinp/go"
	"reflect"
	"strings"
	"time"
)

// AwsAWSEC2Foundation - Model AWSEC2Foundation(/api/v1/AWS/AWSEC2Foundation)
/*
AWSEC2Foundation(locator, site, blueprint, id_map, located_at, built_at, updated, created, foundation_ptr, awsec2_instance_id)
*/
type AwsAWSEC2Foundation struct {
	cinp.BaseObject
	cinp             cinp.CInPClient `json:"-"`
	Locator          *string         `json:"locator,omitempty"`
	Site             *string         `json:"site,omitempty"`
	Blueprint        *string         `json:"blueprint,omitempty"`
	IDMap            *string         `json:"id_map,omitempty"`
	LocatedAt        *time.Time      `json:"located_at,omitempty"`
	BuiltAt          *time.Time      `json:"built_at,omitempty"`
	Updated          *time.Time      `json:"updated,omitempty"`
	Created          *time.Time      `json:"created,omitempty"`
	Awsec2InstanceID *string         `json:"awsec2_instance_id,omitempty"`
	State            *string         `json:"state,omitempty"`
	Type             *string         `json:"type,omitempty"`
	ClassList        *string         `json:"class_list,omitempty"`
}

// AwsAWSEC2FoundationNew - Make a new object of Model AWSEC2Foundation
func (service *Contractor) AwsAWSEC2FoundationNew() *AwsAWSEC2Foundation {
	return &AwsAWSEC2Foundation{cinp: service.cinp}
}

// AwsAWSEC2FoundationNewWithID - Make a new object of Model AWSEC2Foundation
func (service *Contractor) AwsAWSEC2FoundationNewWithID(id string) *AwsAWSEC2Foundation {
	result := AwsAWSEC2Foundation{cinp: service.cinp}
	result.SetURI("/api/v1/AWS/AWSEC2Foundation:" + id + ":")
	return &result
}

// AwsAWSEC2FoundationGet - Get function for Model AWSEC2Foundation
func (service *Contractor) AwsAWSEC2FoundationGet(ctx context.Context, id string) (*AwsAWSEC2Foundation, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/AWS/AWSEC2Foundation:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*AwsAWSEC2Foundation)
	result.cinp = service.cinp

	return result, nil
}

// AwsAWSEC2FoundationGetURI - Get function for Model AWSEC2Foundation vi URI
func (service *Contractor) AwsAWSEC2FoundationGetURI(ctx context.Context, uri string) (*AwsAWSEC2Foundation, error) {
	if !strings.HasPrefix(uri, "/api/v1/AWS/AWSEC2Foundation:") {
		return nil, fmt.Errorf("URI is not for a 'AwsAWSEC2Foundation'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*AwsAWSEC2Foundation)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model AWSEC2Foundation
func (object *AwsAWSEC2Foundation) Create(ctx context.Context) (*AwsAWSEC2Foundation, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/AWS/AWSEC2Foundation", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*AwsAWSEC2Foundation), nil
}

// Update - Update function for Model AWSEC2Foundation
func (object *AwsAWSEC2Foundation) Update(ctx context.Context) (*AwsAWSEC2Foundation, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*AwsAWSEC2Foundation), nil
}

// Delete - Delete function for Model AWSEC2Foundation
func (object *AwsAWSEC2Foundation) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// AwsAWSEC2FoundationListFilters - Return a slice of valid filter names AWSEC2Foundation
func (service *Contractor) AwsAWSEC2FoundationListFilters() [1]string {
	return [1]string{"site"}
}

// AwsAWSEC2FoundationList - List function for Model AWSEC2Foundation
func (service *Contractor) AwsAWSEC2FoundationList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *AwsAWSEC2Foundation, error) {
	if filterName != "" {
		for _, item := range service.AwsAWSEC2FoundationListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/AWS/AWSEC2Foundation", reflect.TypeOf(AwsAWSEC2Foundation{}), filterName, filterValues, 50)
	out := make(chan *AwsAWSEC2Foundation)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*AwsAWSEC2Foundation).cinp = service.cinp
			out <- (*v).(*AwsAWSEC2Foundation)
		}
	}()
	return out, nil
}

func registerAWS(cinp cinp.CInPClient) {
	cinp.RegisterType("/api/v1/AWS/AWSEC2Foundation", reflect.TypeOf((*AwsAWSEC2Foundation)(nil)).Elem())
}
