/*Package contractor(version: "0.1") - Automatically generated by cinp-codegen from /api/v1/Proxmox/ at 2020-08-13T22:23:45.513882
 */
package contractor

import (
	"reflect"
	"time"
	cinp "github.com/cinp/go"
)

//ProxmoxProxmoxComplex - Model ProxmoxComplex(/api/v1/Proxmox/ProxmoxComplex)
/*
ProxmoxComplex(name, site, description, built_percentage, updated, created, complex_ptr, proxmox_username, proxmox_password)
 */
type ProxmoxProxmoxComplex struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Name string `json:"name"`
	Site string `json:"site"`
	Description string `json:"description"`
	BuiltPercentage int `json:"built_percentage"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
	ProxmoxUsername string `json:"proxmox_username"`
	ProxmoxPassword string `json:"proxmox_password"`
	Members []string `json:"members"`
	State string `json:"state"`
	Type string `json:"type"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *ProxmoxProxmoxComplex) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"name": object.Name,
			"site": object.Site,
			"description": object.Description,
			"built_percentage": object.BuiltPercentage,
			"proxmox_username": object.ProxmoxUsername,
			"proxmox_password": object.ProxmoxPassword,
		}
	}
	return &map[string]interface{}{ 
		"site": object.Site,
		"description": object.Description,
		"built_percentage": object.BuiltPercentage,
		"proxmox_username": object.ProxmoxUsername,
		"proxmox_password": object.ProxmoxPassword,
	}
}

// ProxmoxProxmoxComplexNew - Make a new object of Model ProxmoxComplex
func (service *Contractor) ProxmoxProxmoxComplexNew() *ProxmoxProxmoxComplex {
	return &ProxmoxProxmoxComplex{cinp: service.cinp}
}

// ProxmoxProxmoxComplexNewWithID - Make a new object of Model ProxmoxComplex
func (service *Contractor) ProxmoxProxmoxComplexNewWithID(id string) *ProxmoxProxmoxComplex {
	result := ProxmoxProxmoxComplex{cinp: service.cinp}
	result.SetID("/api/v1/Proxmox/ProxmoxComplex:"+id+":")
	return &result
}

// ProxmoxProxmoxComplexGet - Get function for Model ProxmoxComplex
func (service *Contractor) ProxmoxProxmoxComplexGet(id string) (*ProxmoxProxmoxComplex, error) {
	object, err := service.cinp.Get("/api/v1/Proxmox/ProxmoxComplex:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*ProxmoxProxmoxComplex)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model ProxmoxComplex
func (object *ProxmoxProxmoxComplex) Create() error {
	if err := object.cinp.Create("/api/v1/Proxmox/ProxmoxComplex", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model ProxmoxComplex
func (object *ProxmoxProxmoxComplex) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model ProxmoxComplex
func (object *ProxmoxProxmoxComplex) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// ProxmoxProxmoxComplexList - List function for Model ProxmoxComplex
func (service *Contractor) ProxmoxProxmoxComplexList(filterName string, filterValues map[string]interface{}) <-chan *ProxmoxProxmoxComplex {
	in := service.cinp.ListObjects("/api/v1/Proxmox/ProxmoxComplex", reflect.TypeOf(ProxmoxProxmoxComplex{}), filterName, filterValues, 50)
	out := make(chan *ProxmoxProxmoxComplex)
	go func() {
		defer close(out)
		for v := range in {
			v.(*ProxmoxProxmoxComplex).cinp = service.cinp
			out <- v.(*ProxmoxProxmoxComplex)
		}
	}()
	return out
}



//ProxmoxProxmoxFoundation - Model ProxmoxFoundation(/api/v1/Proxmox/ProxmoxFoundation)
/*
ProxmoxFoundation(locator, site, blueprint, id_map, located_at, built_at, updated, created, foundation_ptr, proxmox_complex, proxmox_vmid)
 */
type ProxmoxProxmoxFoundation struct {
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
	ProxmoxComplex string `json:"proxmox_complex"`
	ProxmoxVmid int `json:"proxmox_vmid"`
	State string `json:"state"`
	Type string `json:"type"`
	ClassList string `json:"class_list"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *ProxmoxProxmoxFoundation) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"locator": object.Locator,
			"site": object.Site,
			"blueprint": object.Blueprint,
			"id_map": object.IDMap,
			"proxmox_complex": object.ProxmoxComplex,
		}
	}
	return &map[string]interface{}{ 
		"site": object.Site,
		"blueprint": object.Blueprint,
		"id_map": object.IDMap,
		"proxmox_complex": object.ProxmoxComplex,
	}
}

// ProxmoxProxmoxFoundationNew - Make a new object of Model ProxmoxFoundation
func (service *Contractor) ProxmoxProxmoxFoundationNew() *ProxmoxProxmoxFoundation {
	return &ProxmoxProxmoxFoundation{cinp: service.cinp}
}

// ProxmoxProxmoxFoundationNewWithID - Make a new object of Model ProxmoxFoundation
func (service *Contractor) ProxmoxProxmoxFoundationNewWithID(id string) *ProxmoxProxmoxFoundation {
	result := ProxmoxProxmoxFoundation{cinp: service.cinp}
	result.SetID("/api/v1/Proxmox/ProxmoxFoundation:"+id+":")
	return &result
}

// ProxmoxProxmoxFoundationGet - Get function for Model ProxmoxFoundation
func (service *Contractor) ProxmoxProxmoxFoundationGet(id string) (*ProxmoxProxmoxFoundation, error) {
	object, err := service.cinp.Get("/api/v1/Proxmox/ProxmoxFoundation:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*ProxmoxProxmoxFoundation)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model ProxmoxFoundation
func (object *ProxmoxProxmoxFoundation) Create() error {
	if err := object.cinp.Create("/api/v1/Proxmox/ProxmoxFoundation", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model ProxmoxFoundation
func (object *ProxmoxProxmoxFoundation) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model ProxmoxFoundation
func (object *ProxmoxProxmoxFoundation) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// ProxmoxProxmoxFoundationList - List function for Model ProxmoxFoundation
func (service *Contractor) ProxmoxProxmoxFoundationList(filterName string, filterValues map[string]interface{}) <-chan *ProxmoxProxmoxFoundation {
	in := service.cinp.ListObjects("/api/v1/Proxmox/ProxmoxFoundation", reflect.TypeOf(ProxmoxProxmoxFoundation{}), filterName, filterValues, 50)
	out := make(chan *ProxmoxProxmoxFoundation)
	go func() {
		defer close(out)
		for v := range in {
			v.(*ProxmoxProxmoxFoundation).cinp = service.cinp
			out <- v.(*ProxmoxProxmoxFoundation)
		}
	}()
	return out
}


func registerProxmox(cinp *cinp.CInP) { 
	cinp.RegisterType("/api/v1/Proxmox/ProxmoxComplex", reflect.TypeOf((*ProxmoxProxmoxComplex)(nil)).Elem())
	cinp.RegisterType("/api/v1/Proxmox/ProxmoxFoundation", reflect.TypeOf((*ProxmoxProxmoxFoundation)(nil)).Elem())
}