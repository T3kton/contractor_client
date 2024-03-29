// Package contractor - (version: "0.1") - Automatically generated by cinp-codegen from /api/v1/Docker/ at 2024-03-24T03:21:24.739985
package contractor

import (
	cinp "github.com/cinp/go"
	"fmt"
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
	cinp *cinp.CInP
	Name string `json:"name"`
	Site string `json:"site"`
	Description string `json:"description"`
	BuiltPercentage int `json:"built_percentage"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
	Members []string `json:"members"`
	State string `json:"state"`
	Type string `json:"type"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *DockerDockerComplex) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"name": object.Name,
			"site": object.Site,
			"description": object.Description,
			"built_percentage": object.BuiltPercentage,
			"members": object.Members,
		}
	}
	return &map[string]interface{}{ 
		"site": object.Site,
		"description": object.Description,
		"built_percentage": object.BuiltPercentage,
		"members": object.Members,
	}
}

// DockerDockerComplexNew - Make a new object of Model DockerComplex
func (service *Contractor) DockerDockerComplexNew() *DockerDockerComplex {
	return &DockerDockerComplex{cinp: service.cinp}
}

// DockerDockerComplexNewWithID - Make a new object of Model DockerComplex
func (service *Contractor) DockerDockerComplexNewWithID(id string) *DockerDockerComplex {
	result := DockerDockerComplex{cinp: service.cinp}
	result.SetID("/api/v1/Docker/DockerComplex:" + id + ":")
	return &result
}

// DockerDockerComplexGet - Get function for Model DockerComplex
func (service *Contractor) DockerDockerComplexGet(id string) (*DockerDockerComplex, error) {
	object, err := service.cinp.Get("/api/v1/Docker/DockerComplex:" + id + ":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*DockerDockerComplex)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model DockerComplex
func (object *DockerDockerComplex) Create() error {
	if err := object.cinp.Create("/api/v1/Docker/DockerComplex", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model DockerComplex
func (object *DockerDockerComplex) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model DockerComplex
func (object *DockerDockerComplex) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// DockerDockerComplexListFilters - Return a slice of valid filter names DockerComplex
func (service *Contractor) DockerDockerComplexListFilters() [0]string {
  return [0]string{  }
}

// DockerDockerComplexList - List function for Model DockerComplex
func (service *Contractor) DockerDockerComplexList(filterName string, filterValues map[string]interface{}) (<-chan *DockerDockerComplex, error) {
	if filterName != "" {
		for _, item := range service.DockerDockerComplexListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("Filter '%s' is invalid", filterName)
	}
	good:

	in := service.cinp.ListObjects("/api/v1/Docker/DockerComplex", reflect.TypeOf(DockerDockerComplex{}), filterName, filterValues, 50)
	out := make(chan *DockerDockerComplex)
	go func() {
		defer close(out)
		for v := range in {
			v.(*DockerDockerComplex).cinp = service.cinp
			out <- v.(*DockerDockerComplex)
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
	cinp *cinp.CInP
	Locator string `json:"locator"`
	Site string `json:"site"`
	Blueprint string `json:"blueprint"`
	IDMap string `json:"id_map"`
	LocatedAt time.Time `json:"located_at"`
	BuiltAt time.Time `json:"built_at"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
	DockerComplex string `json:"docker_complex"`
	DockerID string `json:"docker_id"`
	State string `json:"state"`
	Type string `json:"type"`
	ClassList string `json:"class_list"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *DockerDockerFoundation) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"locator": object.Locator,
			"site": object.Site,
			"blueprint": object.Blueprint,
			"id_map": object.IDMap,
			"docker_complex": object.DockerComplex,
		}
	}
	return &map[string]interface{}{ 
		"site": object.Site,
		"blueprint": object.Blueprint,
		"id_map": object.IDMap,
		"docker_complex": object.DockerComplex,
	}
}

// DockerDockerFoundationNew - Make a new object of Model DockerFoundation
func (service *Contractor) DockerDockerFoundationNew() *DockerDockerFoundation {
	return &DockerDockerFoundation{cinp: service.cinp}
}

// DockerDockerFoundationNewWithID - Make a new object of Model DockerFoundation
func (service *Contractor) DockerDockerFoundationNewWithID(id string) *DockerDockerFoundation {
	result := DockerDockerFoundation{cinp: service.cinp}
	result.SetID("/api/v1/Docker/DockerFoundation:" + id + ":")
	return &result
}

// DockerDockerFoundationGet - Get function for Model DockerFoundation
func (service *Contractor) DockerDockerFoundationGet(id string) (*DockerDockerFoundation, error) {
	object, err := service.cinp.Get("/api/v1/Docker/DockerFoundation:" + id + ":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*DockerDockerFoundation)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model DockerFoundation
func (object *DockerDockerFoundation) Create() error {
	if err := object.cinp.Create("/api/v1/Docker/DockerFoundation", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model DockerFoundation
func (object *DockerDockerFoundation) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model DockerFoundation
func (object *DockerDockerFoundation) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// DockerDockerFoundationListFilters - Return a slice of valid filter names DockerFoundation
func (service *Contractor) DockerDockerFoundationListFilters() [1]string {
  return [1]string{ "site" }
}

// DockerDockerFoundationList - List function for Model DockerFoundation
func (service *Contractor) DockerDockerFoundationList(filterName string, filterValues map[string]interface{}) (<-chan *DockerDockerFoundation, error) {
	if filterName != "" {
		for _, item := range service.DockerDockerFoundationListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("Filter '%s' is invalid", filterName)
	}
	good:

	in := service.cinp.ListObjects("/api/v1/Docker/DockerFoundation", reflect.TypeOf(DockerDockerFoundation{}), filterName, filterValues, 50)
	out := make(chan *DockerDockerFoundation)
	go func() {
		defer close(out)
		for v := range in {
			v.(*DockerDockerFoundation).cinp = service.cinp
			out <- v.(*DockerDockerFoundation)
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
	cinp *cinp.CInP
	Complex string `json:"complex"`
	Port int `json:"port"`
	AddressOffset int `json:"address_offset"`
	Foundation string `json:"foundation"`
	FoundationIndex int `json:"foundation_index"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
	ID int `json:"id"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *DockerDockerPort) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"complex": object.Complex,
			"port": object.Port,
			"address_offset": object.AddressOffset,
			"foundation": object.Foundation,
			"foundation_index": object.FoundationIndex,
		}
	}
	return &map[string]interface{}{ 
		"complex": object.Complex,
		"port": object.Port,
		"address_offset": object.AddressOffset,
		"foundation": object.Foundation,
		"foundation_index": object.FoundationIndex,
	}
}

// DockerDockerPortNew - Make a new object of Model DockerPort
func (service *Contractor) DockerDockerPortNew() *DockerDockerPort {
	return &DockerDockerPort{cinp: service.cinp}
}

// DockerDockerPortNewWithID - Make a new object of Model DockerPort
func (service *Contractor) DockerDockerPortNewWithID(id int) *DockerDockerPort {
	result := DockerDockerPort{cinp: service.cinp}
	result.SetID("/api/v1/Docker/DockerPort:" + strconv.FormatInt(int64(id), 10) + ":")
	return &result
}

// DockerDockerPortGet - Get function for Model DockerPort
func (service *Contractor) DockerDockerPortGet(id int) (*DockerDockerPort, error) {
	object, err := service.cinp.Get("/api/v1/Docker/DockerPort:" + strconv.FormatInt(int64(id), 10) + ":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*DockerDockerPort)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model DockerPort
func (object *DockerDockerPort) Create() error {
	if err := object.cinp.Create("/api/v1/Docker/DockerPort", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model DockerPort
func (object *DockerDockerPort) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model DockerPort
func (object *DockerDockerPort) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// DockerDockerPortListFilters - Return a slice of valid filter names DockerPort
func (service *Contractor) DockerDockerPortListFilters() [0]string {
  return [0]string{  }
}

// DockerDockerPortList - List function for Model DockerPort
func (service *Contractor) DockerDockerPortList(filterName string, filterValues map[string]interface{}) (<-chan *DockerDockerPort, error) {
	if filterName != "" {
		for _, item := range service.DockerDockerPortListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("Filter '%s' is invalid", filterName)
	}
	good:

	in := service.cinp.ListObjects("/api/v1/Docker/DockerPort", reflect.TypeOf(DockerDockerPort{}), filterName, filterValues, 50)
	out := make(chan *DockerDockerPort)
	go func() {
		defer close(out)
		for v := range in {
			v.(*DockerDockerPort).cinp = service.cinp
			out <- v.(*DockerDockerPort)
		}
	}()
	return out, nil
}

func registerDocker(cinp *cinp.CInP) {
	cinp.RegisterType("/api/v1/Docker/DockerComplex", reflect.TypeOf((*DockerDockerComplex)(nil)).Elem())
	cinp.RegisterType("/api/v1/Docker/DockerFoundation", reflect.TypeOf((*DockerDockerFoundation)(nil)).Elem())
	cinp.RegisterType("/api/v1/Docker/DockerPort", reflect.TypeOf((*DockerDockerPort)(nil)).Elem())
}
