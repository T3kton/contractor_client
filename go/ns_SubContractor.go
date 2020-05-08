/*Package contractor(version: "0.1") - Automatically generated by cinp-codegen from /api/v1/SubContractor/ at 2020-05-08T12:31:05.921360
 */
package contractor

import (
	"reflect"
	cinp "github.com/cinp/go"
)

//SubcontractorDispatch - Model Dispatch(/api/v1/SubContractor/Dispatch)
/*

 */
type SubcontractorDispatch struct {
	cinp.BaseObject
	cinp *cinp.CInP
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *SubcontractorDispatch) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
		}
	}
	return &map[string]interface{}{ 
	}
}

// SubcontractorDispatchNew - Make a new object of Model Dispatch
func (service *Contractor) SubcontractorDispatchNew() *SubcontractorDispatch {
	return &SubcontractorDispatch{cinp: service.cinp}
}

// SubcontractorDispatchNewWithID - Make a new object of Model Dispatch
func (service *Contractor) SubcontractorDispatchNewWithID(id string) *SubcontractorDispatch {
	result := SubcontractorDispatch{cinp: service.cinp}
	result.SetID("/api/v1/SubContractor/Dispatch:"+id+":")
	return &result
}


// SubcontractorDispatchCallGetJobs calls getJobsNoneNoneNone
func (service *Contractor) SubcontractorDispatchCallGetJobs(site string, module_list []string, max_jobs int) ([]map[string]interface{}, error) {
	args := map[string]interface{}{
		"site": site,
		"module_list": module_list,
		"max_jobs": max_jobs,
	}
	uri := "/api/v1/SubContractor/Dispatch(getJobs)"

	result := []map[string]interface{}{}

	if err := service.cinp.Call(uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// SubcontractorDispatchCallJobResults calls jobResultsNoneNoneNone
func (service *Contractor) SubcontractorDispatchCallJobResults(job_id int, cookie string, data map[string]interface{}) (string, error) {
	args := map[string]interface{}{
		"job_id": job_id,
		"cookie": cookie,
		"data": data,
	}
	uri := "/api/v1/SubContractor/Dispatch(jobResults)"

	result := ""

	if err := service.cinp.Call(uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// SubcontractorDispatchCallJobError calls jobErrorNoneNoneNone
func (service *Contractor) SubcontractorDispatchCallJobError(job_id int, cookie string, msg string) (error) {
	args := map[string]interface{}{
		"job_id": job_id,
		"cookie": cookie,
		"msg": msg,
	}
	uri := "/api/v1/SubContractor/Dispatch(jobError)"

	result := ""

	if err := service.cinp.Call(uri, &args, &result); err != nil {
		return err
	}

	return nil
}


//SubcontractorDHCPd - Model DHCPd(/api/v1/SubContractor/DHCPd)
/*

 */
type SubcontractorDHCPd struct {
	cinp.BaseObject
	cinp *cinp.CInP
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *SubcontractorDHCPd) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
		}
	}
	return &map[string]interface{}{ 
	}
}

// SubcontractorDHCPdNew - Make a new object of Model DHCPd
func (service *Contractor) SubcontractorDHCPdNew() *SubcontractorDHCPd {
	return &SubcontractorDHCPd{cinp: service.cinp}
}

// SubcontractorDHCPdNewWithID - Make a new object of Model DHCPd
func (service *Contractor) SubcontractorDHCPdNewWithID(id string) *SubcontractorDHCPd {
	result := SubcontractorDHCPd{cinp: service.cinp}
	result.SetID("/api/v1/SubContractor/DHCPd:"+id+":")
	return &result
}


// SubcontractorDHCPdCallGetDynamicPools calls getDynamicPoolsNone
func (service *Contractor) SubcontractorDHCPdCallGetDynamicPools(site string) ([]map[string]interface{}, error) {
	args := map[string]interface{}{
		"site": site,
	}
	uri := "/api/v1/SubContractor/DHCPd(getDynamicPools)"

	result := []map[string]interface{}{}

	if err := service.cinp.Call(uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// SubcontractorDHCPdCallGetStaticPools calls getStaticPoolsNone
func (service *Contractor) SubcontractorDHCPdCallGetStaticPools(site string) (map[string]interface{}, error) {
	args := map[string]interface{}{
		"site": site,
	}
	uri := "/api/v1/SubContractor/DHCPd(getStaticPools)"

	result := map[string]interface{}{}

	if err := service.cinp.Call(uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func registerSubContractor(cinp *cinp.CInP) { 
	cinp.RegisterType("/api/v1/SubContractor/Dispatch", reflect.TypeOf((*SubcontractorDispatch)(nil)).Elem())
	cinp.RegisterType("/api/v1/SubContractor/DHCPd", reflect.TypeOf((*SubcontractorDHCPd)(nil)).Elem())
}