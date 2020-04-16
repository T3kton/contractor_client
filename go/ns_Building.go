/*Package contractor(version: "0.1") - Automatically generated by cinp-codegen from /api/v1/Building/ at 2020-04-16T20:18:01.673062
 */
package contractor

import (
	"reflect"
	"time"
	cinp "github.com/cinp/go"
)

//BuildingFoundation - Model Foundation(/api/v1/Building/Foundation)
/*
Foundation(locator, site, blueprint, id_map, located_at, built_at, updated, created)
 */
type BuildingFoundation struct {
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
	State string `json:"state"`
	Type string `json:"type"`
	ClassList string `json:"class_list"`
	Structure string `json:"structure"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *BuildingFoundation) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"locator": object.Locator,
			"site": object.Site,
			"blueprint": object.Blueprint,
			"id_map": object.IDMap,
		}
	}
	return &map[string]interface{}{ 
		"site": object.Site,
		"blueprint": object.Blueprint,
		"id_map": object.IDMap,
	}
}

// BuildingFoundationNew - Make a new object of Model Foundation
func (service *Contractor) BuildingFoundationNew() *BuildingFoundation {
	return &BuildingFoundation{cinp: service.cinp}
}

// BuildingFoundationNewWithID - Make a new object of Model Foundation
func (service *Contractor) BuildingFoundationNewWithID(id string) *BuildingFoundation {
	result := BuildingFoundation{cinp: service.cinp}
	result.SetID("/api/v1/Building/Foundation:"+id+":")
	return &result
}

// BuildingFoundationGet - Get function for Model Foundation
func (service *Contractor) BuildingFoundationGet(id string) (*BuildingFoundation, error) {
	object, err := service.cinp.Get("/api/v1/Building/Foundation:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*BuildingFoundation)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model Foundation
func (object *BuildingFoundation) Create() error {
	if err := object.cinp.Create("/api/v1/Building/Foundation", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model Foundation
func (object *BuildingFoundation) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model Foundation
func (object *BuildingFoundation) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// BuildingFoundationList - List function for Model Foundation
func (service *Contractor) BuildingFoundationList(filterName string, filterValues map[string]interface{}) <-chan *BuildingFoundation {
	in := service.cinp.ListObjects("/api/v1/Building/Foundation", reflect.TypeOf(BuildingFoundation{}), filterName, filterValues, 50)
	out := make(chan *BuildingFoundation)
	go func() {
		defer close(out)
		for v := range in {
			out <- v.(*BuildingFoundation)
		}
	}()
	return out
}

// CallSetIdMap calls setIdMapNone
func (object *BuildingFoundation) CallSetIdMap(id_map map[string]interface{}) (string, error) {
	args := map[string]interface{}{
		"id_map": id_map,
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return "", err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Building/Foundation(setIdMap)", ids)
	if err != nil {
		return "", err
	}

	result := ""

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// CallDoCreate calls doCreate
func (object *BuildingFoundation) CallDoCreate() (int, error) {
	args := map[string]interface{}{
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return 0, err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Building/Foundation(doCreate)", ids)
	if err != nil {
		return 0, err
	}

	result := 0

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return 0, err
	}

	return result, nil
}

// CallDoDestroy calls doDestroy
func (object *BuildingFoundation) CallDoDestroy() (int, error) {
	args := map[string]interface{}{
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return 0, err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Building/Foundation(doDestroy)", ids)
	if err != nil {
		return 0, err
	}

	result := 0

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return 0, err
	}

	return result, nil
}

// CallDoJob calls doJobNone
func (object *BuildingFoundation) CallDoJob(name string) (int, error) {
	args := map[string]interface{}{
		"name": name,
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return 0, err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Building/Foundation(doJob)", ids)
	if err != nil {
		return 0, err
	}

	result := 0

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return 0, err
	}

	return result, nil
}

// CallGetJob calls getJob
func (object *BuildingFoundation) CallGetJob() (string, error) {
	args := map[string]interface{}{
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return "", err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Building/Foundation(getJob)", ids)
	if err != nil {
		return "", err
	}

	result := ""

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// BuildingFoundationCallGetFoundationTypes calls getFoundationTypes
func (service *Contractor) BuildingFoundationCallGetFoundationTypes() ([]string, error) {
	args := map[string]interface{}{
	}
	uri := "/api/v1/Building/Foundation(getFoundationTypes)"

	result := []string{}

	if err := service.cinp.Call(uri, &args, &result); err != nil {
		return []string{}, err
	}

	return result, nil
}

// CallGetConfig calls getConfig
func (object *BuildingFoundation) CallGetConfig() (map[string]interface{}, error) {
	args := map[string]interface{}{
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Building/Foundation(getConfig)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CallGetInterfaceList calls getInterfaceList
func (object *BuildingFoundation) CallGetInterfaceList() ([]map[string]interface{}, error) {
	args := map[string]interface{}{
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Building/Foundation(getInterfaceList)", ids)
	if err != nil {
		return nil, err
	}

	result := []map[string]interface{}{}

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}


//BuildingStructure - Model Structure(/api/v1/Building/Structure)
/*
Structure(id, hostname, site, networked_ptr, blueprint, foundation, config_uuid, config_values, built_at, updated, created)
 */
type BuildingStructure struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Hostname string `json:"hostname"`
	Site string `json:"site"`
	Blueprint string `json:"blueprint"`
	Foundation string `json:"foundation"`
	ConfigUUID string `json:"config_uuid"`
	ConfigValues map[string]interface{} `json:"config_values"`
	BuiltAt time.Time `json:"built_at"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
	State string `json:"state"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *BuildingStructure) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"hostname": object.Hostname,
			"site": object.Site,
			"blueprint": object.Blueprint,
			"foundation": object.Foundation,
			"config_values": object.ConfigValues,
		}
	}
	return &map[string]interface{}{ 
		"hostname": object.Hostname,
		"site": object.Site,
		"blueprint": object.Blueprint,
		"foundation": object.Foundation,
		"config_values": object.ConfigValues,
	}
}

// BuildingStructureNew - Make a new object of Model Structure
func (service *Contractor) BuildingStructureNew() *BuildingStructure {
	return &BuildingStructure{cinp: service.cinp}
}

// BuildingStructureNewWithID - Make a new object of Model Structure
func (service *Contractor) BuildingStructureNewWithID(id string) *BuildingStructure {
	result := BuildingStructure{cinp: service.cinp}
	result.SetID("/api/v1/Building/Structure:"+id+":")
	return &result
}

// BuildingStructureGet - Get function for Model Structure
func (service *Contractor) BuildingStructureGet(id string) (*BuildingStructure, error) {
	object, err := service.cinp.Get("/api/v1/Building/Structure:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*BuildingStructure)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model Structure
func (object *BuildingStructure) Create() error {
	if err := object.cinp.Create("/api/v1/Building/Structure", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model Structure
func (object *BuildingStructure) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model Structure
func (object *BuildingStructure) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// BuildingStructureList - List function for Model Structure
func (service *Contractor) BuildingStructureList(filterName string, filterValues map[string]interface{}) <-chan *BuildingStructure {
	in := service.cinp.ListObjects("/api/v1/Building/Structure", reflect.TypeOf(BuildingStructure{}), filterName, filterValues, 50)
	out := make(chan *BuildingStructure)
	go func() {
		defer close(out)
		for v := range in {
			out <- v.(*BuildingStructure)
		}
	}()
	return out
}

// CallDoCreate calls doCreate
func (object *BuildingStructure) CallDoCreate() (int, error) {
	args := map[string]interface{}{
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return 0, err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Building/Structure(doCreate)", ids)
	if err != nil {
		return 0, err
	}

	result := 0

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return 0, err
	}

	return result, nil
}

// CallDoDestroy calls doDestroy
func (object *BuildingStructure) CallDoDestroy() (int, error) {
	args := map[string]interface{}{
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return 0, err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Building/Structure(doDestroy)", ids)
	if err != nil {
		return 0, err
	}

	result := 0

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return 0, err
	}

	return result, nil
}

// CallDoJob calls doJobNone
func (object *BuildingStructure) CallDoJob(name string) (int, error) {
	args := map[string]interface{}{
		"name": name,
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return 0, err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Building/Structure(doJob)", ids)
	if err != nil {
		return 0, err
	}

	result := 0

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return 0, err
	}

	return result, nil
}

// CallGetJob calls getJob
func (object *BuildingStructure) CallGetJob() (string, error) {
	args := map[string]interface{}{
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return "", err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Building/Structure(getJob)", ids)
	if err != nil {
		return "", err
	}

	result := ""

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// CallGetConfig calls getConfig
func (object *BuildingStructure) CallGetConfig() (map[string]interface{}, error) {
	args := map[string]interface{}{
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Building/Structure(getConfig)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CallUpdateConfig calls updateConfigNone
func (object *BuildingStructure) CallUpdateConfig(config_value_map map[string]interface{}) (map[string]interface{}, error) {
	args := map[string]interface{}{
		"config_value_map": config_value_map,
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Building/Structure(updateConfig)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}


//BuildingComplex - Model Complex(/api/v1/Building/Complex)
/*
Complex(name, site, description, built_percentage, updated, created)
 */
type BuildingComplex struct {
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
func (object *BuildingComplex) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"name": object.Name,
			"site": object.Site,
			"description": object.Description,
			"built_percentage": object.BuiltPercentage,
		}
	}
	return &map[string]interface{}{ 
		"site": object.Site,
		"description": object.Description,
		"built_percentage": object.BuiltPercentage,
	}
}

// BuildingComplexNew - Make a new object of Model Complex
func (service *Contractor) BuildingComplexNew() *BuildingComplex {
	return &BuildingComplex{cinp: service.cinp}
}

// BuildingComplexNewWithID - Make a new object of Model Complex
func (service *Contractor) BuildingComplexNewWithID(id string) *BuildingComplex {
	result := BuildingComplex{cinp: service.cinp}
	result.SetID("/api/v1/Building/Complex:"+id+":")
	return &result
}

// BuildingComplexGet - Get function for Model Complex
func (service *Contractor) BuildingComplexGet(id string) (*BuildingComplex, error) {
	object, err := service.cinp.Get("/api/v1/Building/Complex:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*BuildingComplex)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model Complex
func (object *BuildingComplex) Create() error {
	if err := object.cinp.Create("/api/v1/Building/Complex", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model Complex
func (object *BuildingComplex) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model Complex
func (object *BuildingComplex) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// BuildingComplexList - List function for Model Complex
func (service *Contractor) BuildingComplexList(filterName string, filterValues map[string]interface{}) <-chan *BuildingComplex {
	in := service.cinp.ListObjects("/api/v1/Building/Complex", reflect.TypeOf(BuildingComplex{}), filterName, filterValues, 50)
	out := make(chan *BuildingComplex)
	go func() {
		defer close(out)
		for v := range in {
			out <- v.(*BuildingComplex)
		}
	}()
	return out
}

// CallCreateFoundation calls createFoundationNoneNone
func (object *BuildingComplex) CallCreateFoundation(hostname string, interface_name_list []string) (string, error) {
	args := map[string]interface{}{
		"hostname": hostname,
		"interface_name_list": interface_name_list,
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return "", err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Building/Complex(createFoundation)", ids)
	if err != nil {
		return "", err
	}

	result := ""

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}


//BuildingComplexStructure - Model ComplexStructure(/api/v1/Building/ComplexStructure)
/*
ComplexStructure(id, complex, structure, updated, created)
 */
type BuildingComplexStructure struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Complex string `json:"complex"`
	Structure string `json:"structure"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *BuildingComplexStructure) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"complex": object.Complex,
			"structure": object.Structure,
		}
	}
	return &map[string]interface{}{ 
		"complex": object.Complex,
		"structure": object.Structure,
	}
}

// BuildingComplexStructureNew - Make a new object of Model ComplexStructure
func (service *Contractor) BuildingComplexStructureNew() *BuildingComplexStructure {
	return &BuildingComplexStructure{cinp: service.cinp}
}

// BuildingComplexStructureNewWithID - Make a new object of Model ComplexStructure
func (service *Contractor) BuildingComplexStructureNewWithID(id string) *BuildingComplexStructure {
	result := BuildingComplexStructure{cinp: service.cinp}
	result.SetID("/api/v1/Building/ComplexStructure:"+id+":")
	return &result
}

// BuildingComplexStructureGet - Get function for Model ComplexStructure
func (service *Contractor) BuildingComplexStructureGet(id string) (*BuildingComplexStructure, error) {
	object, err := service.cinp.Get("/api/v1/Building/ComplexStructure:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*BuildingComplexStructure)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model ComplexStructure
func (object *BuildingComplexStructure) Create() error {
	if err := object.cinp.Create("/api/v1/Building/ComplexStructure", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model ComplexStructure
func (object *BuildingComplexStructure) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model ComplexStructure
func (object *BuildingComplexStructure) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// BuildingComplexStructureList - List function for Model ComplexStructure
func (service *Contractor) BuildingComplexStructureList(filterName string, filterValues map[string]interface{}) <-chan *BuildingComplexStructure {
	in := service.cinp.ListObjects("/api/v1/Building/ComplexStructure", reflect.TypeOf(BuildingComplexStructure{}), filterName, filterValues, 50)
	out := make(chan *BuildingComplexStructure)
	go func() {
		defer close(out)
		for v := range in {
			out <- v.(*BuildingComplexStructure)
		}
	}()
	return out
}

// CallGetConfig calls getConfig
func (object *BuildingComplexStructure) CallGetConfig() (map[string]interface{}, error) {
	args := map[string]interface{}{
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Building/ComplexStructure(getConfig)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}


//BuildingDependency - Model Dependency(/api/v1/Building/Dependency)
/*
Dependency(id, structure, dependency, foundation, script_structure, link, create_script_name, destroy_script_name, built_at, updated, created)
 */
type BuildingDependency struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Structure string `json:"structure"`
	Dependency string `json:"dependency"`
	Foundation string `json:"foundation"`
	ScriptStructure string `json:"script_structure"`
	Link string `json:"link"`
	CreateScriptName string `json:"create_script_name"`
	DestroyScriptName string `json:"destroy_script_name"`
	BuiltAt time.Time `json:"built_at"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
	State string `json:"state"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *BuildingDependency) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"structure": object.Structure,
			"dependency": object.Dependency,
			"foundation": object.Foundation,
			"script_structure": object.ScriptStructure,
			"link": object.Link,
			"create_script_name": object.CreateScriptName,
			"destroy_script_name": object.DestroyScriptName,
		}
	}
	return &map[string]interface{}{ 
		"structure": object.Structure,
		"dependency": object.Dependency,
		"foundation": object.Foundation,
		"script_structure": object.ScriptStructure,
		"link": object.Link,
		"create_script_name": object.CreateScriptName,
		"destroy_script_name": object.DestroyScriptName,
	}
}

// BuildingDependencyNew - Make a new object of Model Dependency
func (service *Contractor) BuildingDependencyNew() *BuildingDependency {
	return &BuildingDependency{cinp: service.cinp}
}

// BuildingDependencyNewWithID - Make a new object of Model Dependency
func (service *Contractor) BuildingDependencyNewWithID(id string) *BuildingDependency {
	result := BuildingDependency{cinp: service.cinp}
	result.SetID("/api/v1/Building/Dependency:"+id+":")
	return &result
}

// BuildingDependencyGet - Get function for Model Dependency
func (service *Contractor) BuildingDependencyGet(id string) (*BuildingDependency, error) {
	object, err := service.cinp.Get("/api/v1/Building/Dependency:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*BuildingDependency)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model Dependency
func (object *BuildingDependency) Create() error {
	if err := object.cinp.Create("/api/v1/Building/Dependency", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model Dependency
func (object *BuildingDependency) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model Dependency
func (object *BuildingDependency) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// BuildingDependencyList - List function for Model Dependency
func (service *Contractor) BuildingDependencyList(filterName string, filterValues map[string]interface{}) <-chan *BuildingDependency {
	in := service.cinp.ListObjects("/api/v1/Building/Dependency", reflect.TypeOf(BuildingDependency{}), filterName, filterValues, 50)
	out := make(chan *BuildingDependency)
	go func() {
		defer close(out)
		for v := range in {
			out <- v.(*BuildingDependency)
		}
	}()
	return out
}

// CallGetJob calls getJob
func (object *BuildingDependency) CallGetJob() (string, error) {
	args := map[string]interface{}{
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return "", err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Building/Dependency(getJob)", ids)
	if err != nil {
		return "", err
	}

	result := ""

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

func registerBuilding(cinp *cinp.CInP) { 
	cinp.RegisterType("/api/v1/Building/Foundation", reflect.TypeOf((*BuildingFoundation)(nil)).Elem())
	cinp.RegisterType("/api/v1/Building/Structure", reflect.TypeOf((*BuildingStructure)(nil)).Elem())
	cinp.RegisterType("/api/v1/Building/Complex", reflect.TypeOf((*BuildingComplex)(nil)).Elem())
	cinp.RegisterType("/api/v1/Building/ComplexStructure", reflect.TypeOf((*BuildingComplexStructure)(nil)).Elem())
	cinp.RegisterType("/api/v1/Building/Dependency", reflect.TypeOf((*BuildingDependency)(nil)).Elem())
}