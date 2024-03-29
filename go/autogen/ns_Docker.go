// Package contractor - (version: "0.1") - Automatically generated by cinp-codegen from /api/v1/Docker/ at 2024-03-29T15:40:31.896236
package contractor

import (
	"context"
	"fmt"
	cinp "github.com/cinp/go"
	"reflect"
	"strconv"
	"time"
)

// DockerDockerComplex - Model DockerComplex(/api/v1/Docker/DockerComplex)
/*
DockerComplex(name, site, description, built_percentage, updated, created, complex_ptr)
*/
type DockerDockerComplex struct {
	cinp.BaseObject
	cinp            *cinp.CInP `json:"-"`
	Name            *string    `json:"name,omitempty"`
	Site            *string    `json:"site,omitempty"`
	Description     *string    `json:"description,omitempty"`
	BuiltPercentage *int       `json:"built_percentage,omitempty"`
	Updated         *time.Time `json:"updated,omitempty"`
	Created         *time.Time `json:"created,omitempty"`
	Members         *[]string  `json:"members,omitempty"`
	State           *string    `json:"state,omitempty"`
	Type            *string    `json:"type,omitempty"`
}

// DockerDockerComplexNew - Make a new object of Model DockerComplex
func (service *Contractor) DockerDockerComplexNew() *DockerDockerComplex {
	return &DockerDockerComplex{cinp: service.cinp}
}

// DockerDockerComplexNewWithID - Make a new object of Model DockerComplex
func (service *Contractor) DockerDockerComplexNewWithID(id string) *DockerDockerComplex {
	result := DockerDockerComplex{cinp: service.cinp}
	result.SetURI("/api/v1/Docker/DockerComplex:" + id + ":")
	return &result
}

// DockerDockerComplexGet - Get function for Model DockerComplex
func (service *Contractor) DockerDockerComplexGet(ctx context.Context, id string) (*DockerDockerComplex, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Docker/DockerComplex:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*DockerDockerComplex)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model DockerComplex
func (object *DockerDockerComplex) Create(ctx context.Context) (*DockerDockerComplex, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/Docker/DockerComplex", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*DockerDockerComplex), nil
}

// Update - Update function for Model DockerComplex
func (object *DockerDockerComplex) Update(ctx context.Context) (*DockerDockerComplex, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*DockerDockerComplex), nil
}

// Delete - Delete function for Model DockerComplex
func (object *DockerDockerComplex) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// DockerDockerComplexListFilters - Return a slice of valid filter names DockerComplex
func (service *Contractor) DockerDockerComplexListFilters() [0]string {
	return [0]string{}
}

// DockerDockerComplexList - List function for Model DockerComplex
func (service *Contractor) DockerDockerComplexList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *DockerDockerComplex, error) {
	if filterName != "" {
		for _, item := range service.DockerDockerComplexListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("Filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Docker/DockerComplex", reflect.TypeOf(DockerDockerComplex{}), filterName, filterValues, 50)
	out := make(chan *DockerDockerComplex)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*DockerDockerComplex).cinp = service.cinp
			out <- (*v).(*DockerDockerComplex)
		}
	}()
	return out, nil
}

// DockerDockerFoundation - Model DockerFoundation(/api/v1/Docker/DockerFoundation)
/*
DockerFoundation(locator, site, blueprint, id_map, located_at, built_at, updated, created, foundation_ptr, docker_complex, docker_id)
*/
type DockerDockerFoundation struct {
	cinp.BaseObject
	cinp          *cinp.CInP `json:"-"`
	Locator       *string    `json:"locator,omitempty"`
	Site          *string    `json:"site,omitempty"`
	Blueprint     *string    `json:"blueprint,omitempty"`
	IDMap         *string    `json:"id_map,omitempty"`
	LocatedAt     *time.Time `json:"located_at,omitempty"`
	BuiltAt       *time.Time `json:"built_at,omitempty"`
	Updated       *time.Time `json:"updated,omitempty"`
	Created       *time.Time `json:"created,omitempty"`
	DockerComplex *string    `json:"docker_complex,omitempty"`
	DockerID      *string    `json:"docker_id,omitempty"`
	State         *string    `json:"state,omitempty"`
	Type          *string    `json:"type,omitempty"`
	ClassList     *string    `json:"class_list,omitempty"`
}

// DockerDockerFoundationNew - Make a new object of Model DockerFoundation
func (service *Contractor) DockerDockerFoundationNew() *DockerDockerFoundation {
	return &DockerDockerFoundation{cinp: service.cinp}
}

// DockerDockerFoundationNewWithID - Make a new object of Model DockerFoundation
func (service *Contractor) DockerDockerFoundationNewWithID(id string) *DockerDockerFoundation {
	result := DockerDockerFoundation{cinp: service.cinp}
	result.SetURI("/api/v1/Docker/DockerFoundation:" + id + ":")
	return &result
}

// DockerDockerFoundationGet - Get function for Model DockerFoundation
func (service *Contractor) DockerDockerFoundationGet(ctx context.Context, id string) (*DockerDockerFoundation, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Docker/DockerFoundation:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*DockerDockerFoundation)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model DockerFoundation
func (object *DockerDockerFoundation) Create(ctx context.Context) (*DockerDockerFoundation, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/Docker/DockerFoundation", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*DockerDockerFoundation), nil
}

// Update - Update function for Model DockerFoundation
func (object *DockerDockerFoundation) Update(ctx context.Context) (*DockerDockerFoundation, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*DockerDockerFoundation), nil
}

// Delete - Delete function for Model DockerFoundation
func (object *DockerDockerFoundation) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// DockerDockerFoundationListFilters - Return a slice of valid filter names DockerFoundation
func (service *Contractor) DockerDockerFoundationListFilters() [1]string {
	return [1]string{"site"}
}

// DockerDockerFoundationList - List function for Model DockerFoundation
func (service *Contractor) DockerDockerFoundationList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *DockerDockerFoundation, error) {
	if filterName != "" {
		for _, item := range service.DockerDockerFoundationListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("Filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Docker/DockerFoundation", reflect.TypeOf(DockerDockerFoundation{}), filterName, filterValues, 50)
	out := make(chan *DockerDockerFoundation)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*DockerDockerFoundation).cinp = service.cinp
			out <- (*v).(*DockerDockerFoundation)
		}
	}()
	return out, nil
}

// DockerDockerPort - Model DockerPort(/api/v1/Docker/DockerPort)
/*
DockerPort(id, complex, port, address_offset, foundation, foundation_index, updated, created)
*/
type DockerDockerPort struct {
	cinp.BaseObject
	cinp            *cinp.CInP `json:"-"`
	Complex         *string    `json:"complex,omitempty"`
	Port            *int       `json:"port,omitempty"`
	AddressOffset   *int       `json:"address_offset,omitempty"`
	Foundation      *string    `json:"foundation,omitempty"`
	FoundationIndex *int       `json:"foundation_index,omitempty"`
	Updated         *time.Time `json:"updated,omitempty"`
	Created         *time.Time `json:"created,omitempty"`
	ID              *int       `json:"id,omitempty"`
}

// DockerDockerPortNew - Make a new object of Model DockerPort
func (service *Contractor) DockerDockerPortNew() *DockerDockerPort {
	return &DockerDockerPort{cinp: service.cinp}
}

// DockerDockerPortNewWithID - Make a new object of Model DockerPort
func (service *Contractor) DockerDockerPortNewWithID(id int) *DockerDockerPort {
	result := DockerDockerPort{cinp: service.cinp}
	result.SetURI("/api/v1/Docker/DockerPort:" + strconv.FormatInt(int64(id), 10) + ":")
	return &result
}

// DockerDockerPortGet - Get function for Model DockerPort
func (service *Contractor) DockerDockerPortGet(ctx context.Context, id int) (*DockerDockerPort, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Docker/DockerPort:"+strconv.FormatInt(int64(id), 10)+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*DockerDockerPort)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model DockerPort
func (object *DockerDockerPort) Create(ctx context.Context) (*DockerDockerPort, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/Docker/DockerPort", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*DockerDockerPort), nil
}

// Update - Update function for Model DockerPort
func (object *DockerDockerPort) Update(ctx context.Context) (*DockerDockerPort, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*DockerDockerPort), nil
}

// Delete - Delete function for Model DockerPort
func (object *DockerDockerPort) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// DockerDockerPortListFilters - Return a slice of valid filter names DockerPort
func (service *Contractor) DockerDockerPortListFilters() [0]string {
	return [0]string{}
}

// DockerDockerPortList - List function for Model DockerPort
func (service *Contractor) DockerDockerPortList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *DockerDockerPort, error) {
	if filterName != "" {
		for _, item := range service.DockerDockerPortListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("Filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Docker/DockerPort", reflect.TypeOf(DockerDockerPort{}), filterName, filterValues, 50)
	out := make(chan *DockerDockerPort)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*DockerDockerPort).cinp = service.cinp
			out <- (*v).(*DockerDockerPort)
		}
	}()
	return out, nil
}

func registerDocker(cinp *cinp.CInP) {
	cinp.RegisterType("/api/v1/Docker/DockerComplex", reflect.TypeOf((*DockerDockerComplex)(nil)).Elem())
	cinp.RegisterType("/api/v1/Docker/DockerFoundation", reflect.TypeOf((*DockerDockerFoundation)(nil)).Elem())
	cinp.RegisterType("/api/v1/Docker/DockerPort", reflect.TypeOf((*DockerDockerPort)(nil)).Elem())
}
