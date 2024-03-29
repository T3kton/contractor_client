// Package contractor - (version: "0.1") - Automatically generated by cinp-codegen from /api/v1/Records/ at 2024-03-29T15:40:31.896236
package contractor

import (
	"context"
	cinp "github.com/cinp/go"
	"reflect"
)

// RecordsRecorder - Model Recorder(/api/v1/Records/Recorder)
/*

 */
type RecordsRecorder struct {
	cinp.BaseObject
	cinp *cinp.CInP `json:"-"`
}

// RecordsRecorderNew - Make a new object of Model Recorder
func (service *Contractor) RecordsRecorderNew() *RecordsRecorder {
	return &RecordsRecorder{cinp: service.cinp}
}

// RecordsRecorderCallQuery calls query
func (service *Contractor) RecordsRecorderCallQuery(ctx context.Context, Group string, Query string, Fields string, MaxResults int) ([]string, error) {
	args := map[string]interface{}{
		"group":       Group,
		"query":       Query,
		"fields":      Fields,
		"max_results": MaxResults,
	}
	uri := "/api/v1/Records/Recorder(query)"

	result := []string{}

	if err := service.cinp.Call(ctx, uri, &args, &result); err != nil {
		return []string{}, err
	}

	return result, nil
}

// RecordsRecorderCallQueryObjects calls queryObjects
func (service *Contractor) RecordsRecorderCallQueryObjects(ctx context.Context, Group string, Query string, MaxResults int) (string, error) {
	args := map[string]interface{}{
		"group":       Group,
		"query":       Query,
		"max_results": MaxResults,
	}
	uri := "/api/v1/Records/Recorder(queryObjects)"

	result := ""

	if err := service.cinp.Call(ctx, uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}
func registerRecords(cinp *cinp.CInP) {
	cinp.RegisterType("/api/v1/Records/Recorder", reflect.TypeOf((*RecordsRecorder)(nil)).Elem())
}