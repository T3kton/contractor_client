/*Package contractor(version: "0.1") - Automatically generated by cinp-codegen from /api/v1/Packet/ at 2021-09-28T20:47:20.561620
 */
package contractor

import (
	"reflect"
	"time"
	cinp "github.com/cinp/go"
)

//PacketPacketComplex - Model PacketComplex(/api/v1/Packet/PacketComplex)
/*
PacketComplex(name, site, description, built_percentage, updated, created, complex_ptr, packet_auth_token, packet_facility, packet_project)
 */
type PacketPacketComplex struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Name string `json:"name"`
	Site string `json:"site"`
	Description string `json:"description"`
	BuiltPercentage int `json:"built_percentage"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
	PacketAuthToken string `json:"packet_auth_token"`
	PacketFacility string `json:"packet_facility"`
	PacketProject string `json:"packet_project"`
	Members []string `json:"members"`
	State string `json:"state"`
	Type string `json:"type"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *PacketPacketComplex) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"name": object.Name,
			"site": object.Site,
			"description": object.Description,
			"built_percentage": object.BuiltPercentage,
			"packet_auth_token": object.PacketAuthToken,
			"packet_facility": object.PacketFacility,
			"packet_project": object.PacketProject,
			"members": object.Members,
		}
	}
	return &map[string]interface{}{ 
		"site": object.Site,
		"description": object.Description,
		"built_percentage": object.BuiltPercentage,
		"packet_auth_token": object.PacketAuthToken,
		"packet_facility": object.PacketFacility,
		"packet_project": object.PacketProject,
		"members": object.Members,
	}
}

// PacketPacketComplexNew - Make a new object of Model PacketComplex
func (service *Contractor) PacketPacketComplexNew() *PacketPacketComplex {
	return &PacketPacketComplex{cinp: service.cinp}
}

// PacketPacketComplexNewWithID - Make a new object of Model PacketComplex
func (service *Contractor) PacketPacketComplexNewWithID(id string) *PacketPacketComplex {
	result := PacketPacketComplex{cinp: service.cinp}
	result.SetID("/api/v1/Packet/PacketComplex:"+id+":")
	return &result
}

// PacketPacketComplexGet - Get function for Model PacketComplex
func (service *Contractor) PacketPacketComplexGet(id string) (*PacketPacketComplex, error) {
	object, err := service.cinp.Get("/api/v1/Packet/PacketComplex:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*PacketPacketComplex)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model PacketComplex
func (object *PacketPacketComplex) Create() error {
	if err := object.cinp.Create("/api/v1/Packet/PacketComplex", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model PacketComplex
func (object *PacketPacketComplex) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model PacketComplex
func (object *PacketPacketComplex) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// PacketPacketComplexList - List function for Model PacketComplex
func (service *Contractor) PacketPacketComplexList(filterName string, filterValues map[string]interface{}) <-chan *PacketPacketComplex {
	in := service.cinp.ListObjects("/api/v1/Packet/PacketComplex", reflect.TypeOf(PacketPacketComplex{}), filterName, filterValues, 50)
	out := make(chan *PacketPacketComplex)
	go func() {
		defer close(out)
		for v := range in {
			v.(*PacketPacketComplex).cinp = service.cinp
			out <- v.(*PacketPacketComplex)
		}
	}()
	return out
}



//PacketPacketFoundation - Model PacketFoundation(/api/v1/Packet/PacketFoundation)
/*
PacketFoundation(locator, site, blueprint, id_map, located_at, built_at, updated, created, foundation_ptr, packet_complex, packet_uuid)
 */
type PacketPacketFoundation struct {
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
	PacketComplex string `json:"packet_complex"`
	PacketUUID string `json:"packet_uuid"`
	State string `json:"state"`
	Type string `json:"type"`
	ClassList string `json:"class_list"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *PacketPacketFoundation) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"locator": object.Locator,
			"site": object.Site,
			"blueprint": object.Blueprint,
			"id_map": object.IDMap,
			"packet_complex": object.PacketComplex,
		}
	}
	return &map[string]interface{}{ 
		"site": object.Site,
		"blueprint": object.Blueprint,
		"id_map": object.IDMap,
		"packet_complex": object.PacketComplex,
	}
}

// PacketPacketFoundationNew - Make a new object of Model PacketFoundation
func (service *Contractor) PacketPacketFoundationNew() *PacketPacketFoundation {
	return &PacketPacketFoundation{cinp: service.cinp}
}

// PacketPacketFoundationNewWithID - Make a new object of Model PacketFoundation
func (service *Contractor) PacketPacketFoundationNewWithID(id string) *PacketPacketFoundation {
	result := PacketPacketFoundation{cinp: service.cinp}
	result.SetID("/api/v1/Packet/PacketFoundation:"+id+":")
	return &result
}

// PacketPacketFoundationGet - Get function for Model PacketFoundation
func (service *Contractor) PacketPacketFoundationGet(id string) (*PacketPacketFoundation, error) {
	object, err := service.cinp.Get("/api/v1/Packet/PacketFoundation:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*PacketPacketFoundation)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model PacketFoundation
func (object *PacketPacketFoundation) Create() error {
	if err := object.cinp.Create("/api/v1/Packet/PacketFoundation", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model PacketFoundation
func (object *PacketPacketFoundation) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model PacketFoundation
func (object *PacketPacketFoundation) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// PacketPacketFoundationList - List function for Model PacketFoundation
func (service *Contractor) PacketPacketFoundationList(filterName string, filterValues map[string]interface{}) <-chan *PacketPacketFoundation {
	in := service.cinp.ListObjects("/api/v1/Packet/PacketFoundation", reflect.TypeOf(PacketPacketFoundation{}), filterName, filterValues, 50)
	out := make(chan *PacketPacketFoundation)
	go func() {
		defer close(out)
		for v := range in {
			v.(*PacketPacketFoundation).cinp = service.cinp
			out <- v.(*PacketPacketFoundation)
		}
	}()
	return out
}


func registerPacket(cinp *cinp.CInP) { 
	cinp.RegisterType("/api/v1/Packet/PacketComplex", reflect.TypeOf((*PacketPacketComplex)(nil)).Elem())
	cinp.RegisterType("/api/v1/Packet/PacketFoundation", reflect.TypeOf((*PacketPacketFoundation)(nil)).Elem())
}