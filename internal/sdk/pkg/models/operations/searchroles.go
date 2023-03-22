// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"epilotpermissions/internal/sdk/pkg/models/shared"
	"fmt"
	"net/http"
)

type SearchRolesRequest struct {
	Request *shared.RoleSearchInput `request:"mediaType=application/json"`
}

func NewSearchRolesRequest(input interface{}) (*SearchRolesRequest, error) {
	mapInput, ok := input.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("searchRolesRequest: Expected input to be a map[string]interface{}")
	}
	var request *shared.RoleSearchInput
	if mapInput["Request"] != nil {
		requestTmp, err := shared.NewRoleSearchInput(mapInput["Request"])
		if err != nil {
			return nil, err
		}
		request = requestTmp
	}
	out := &SearchRolesRequest{
		Request: request,
	}

	return out, nil
}

// SearchRoles200ApplicationJSON - ok
type SearchRoles200ApplicationJSON struct {
	Hits    *float64      `json:"hits,omitempty"`
	Results []interface{} `json:"results,omitempty"`
}

func NewSearchRoles200ApplicationJSON(input interface{}) (*SearchRoles200ApplicationJSON, error) {
	mapInput, ok := input.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("searchRoles_200ApplicationJSON: Expected input to be a map[string]interface{}")
	}
	hits := new(float64)
	if _, ok = mapInput["hits"]; !ok {
		hits = nil
	} else {
		*hits, ok = mapInput["hits"].(float64)
		if !ok {
			return nil, fmt.Errorf("searchRoles_200ApplicationJSON: unexpected type for Hits. Expected float64 but was %T", mapInput["hits"])
		}
	}
	var results []interface{}
	if _, ok = mapInput["results"]; ok {
		resultsTmp, ok := mapInput["results"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("searchRoles_200ApplicationJSON: unexpected type for Results. Expected []interface{} but was %T", mapInput["results"])
		}
		for _, resultsItemRaw := range resultsTmp {
			var resultsItem interface{}
			pickFirstValidResultsItem := make([]func(interface{}) (interface{}, error), 0)
			pickFirstValidResultsItem = append(pickFirstValidResultsItem, func(sharedUserRoleInput interface{}) (interface{}, error) {
				resultsItemPtr, err := shared.NewUserRole(sharedUserRoleInput)
				if err != nil {
					return nil, err
				}
				sharedUserRoleOutput := *resultsItemPtr
				return sharedUserRoleOutput, nil
			})
			pickFirstValidResultsItem = append(pickFirstValidResultsItem, func(sharedOrgRoleInput interface{}) (interface{}, error) {
				resultsItemPtr1, err := shared.NewOrgRole(sharedOrgRoleInput)
				if err != nil {
					return nil, err
				}
				sharedOrgRoleOutput := *resultsItemPtr1
				return sharedOrgRoleOutput, nil
			})
			pickFirstValidResultsItem = append(pickFirstValidResultsItem, func(sharedShareRoleInput interface{}) (interface{}, error) {
				resultsItemPtr2, err := shared.NewShareRole(sharedShareRoleInput)
				if err != nil {
					return nil, err
				}
				sharedShareRoleOutput := *resultsItemPtr2
				return sharedShareRoleOutput, nil
			})
			pickFirstValidResultsItem = append(pickFirstValidResultsItem, func(sharedPartnerRoleInput interface{}) (interface{}, error) {
				resultsItemPtr3, err := shared.NewPartnerRole(sharedPartnerRoleInput)
				if err != nil {
					return nil, err
				}
				sharedPartnerRoleOutput := *resultsItemPtr3
				return sharedPartnerRoleOutput, nil
			})
			for _, resultsItemChecker := range pickFirstValidResultsItem {
				var resultsItemErr error
				resultsItem, resultsItemErr = resultsItemChecker(resultsItemRaw)
				if resultsItemErr == nil {
					break
				}
			}
			if resultsItem == nil {
				return nil, fmt.Errorf("searchRoles_200ApplicationJSON: unexpected type for ResultsItem. Expected one of shared.UserRole, shared.OrgRole, shared.ShareRole, shared.PartnerRole to be valid. Got %#v", resultsItemRaw)
			}
			results = append(results, resultsItem)
		}
	}
	out := &SearchRoles200ApplicationJSON{
		Hits:    hits,
		Results: results,
	}

	return out, nil
}

type SearchRolesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// ok
	SearchRoles200ApplicationJSONObject *SearchRoles200ApplicationJSON
}

func NewSearchRolesResponse(input interface{}) (*SearchRolesResponse, error) {
	mapInput, ok := input.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("searchRolesResponse: Expected input to be a map[string]interface{}")
	}
	if _, ok = mapInput["ContentType"]; !ok {
		return nil, fmt.Errorf("searchRolesResponse: ContentType is required, but was not found")
	}
	var contentType string
	contentType, ok = mapInput["ContentType"].(string)
	if !ok {
		return nil, fmt.Errorf("searchRolesResponse: unexpected type for ContentType. Expected string but was %T", mapInput["ContentType"])
	}
	if _, ok = mapInput["StatusCode"]; !ok {
		return nil, fmt.Errorf("searchRolesResponse: StatusCode is required, but was not found")
	}
	if _, ok = mapInput["StatusCode"]; !ok {
		return nil, fmt.Errorf("searchRolesResponse: StatusCode is required, but was not found")
	}
	var statusCodeFloat64 float64
	statusCodeFloat64, ok = mapInput["StatusCode"].(float64)
	if !ok {
		return nil, fmt.Errorf("searchRolesResponse: unexpected type for StatusCode. Expected float64 but was %T", mapInput["StatusCode"])
	}
	var statusCode int
	if statusCodeFloat64 != float64(int(statusCodeFloat64)) {
		return nil, fmt.Errorf("searchRolesResponse: unexpected value for integer StatusCode. Got %#v", mapInput["StatusCode"])
	} else {
		statusCode = int(statusCodeFloat64)
	}
	rawResponse := new(http.Response)
	if _, ok = mapInput["RawResponse"]; !ok {
		rawResponse = nil
	} else {
		*rawResponse, ok = mapInput["RawResponse"].(http.Response)
		if !ok {
			return nil, fmt.Errorf("searchRolesResponse: unexpected type for RawResponse. Expected http.Response but was %T", mapInput["RawResponse"])
		}
	}
	var searchRoles200ApplicationJSONObject *SearchRoles200ApplicationJSON
	if mapInput["searchRoles_200ApplicationJSON_object"] != nil {
		searchRoles200ApplicationJSONObjectTmp, err := NewSearchRoles200ApplicationJSON(mapInput["searchRoles_200ApplicationJSON_object"])
		if err != nil {
			return nil, err
		}
		searchRoles200ApplicationJSONObject = searchRoles200ApplicationJSONObjectTmp
	}
	out := &SearchRolesResponse{
		ContentType:                         contentType,
		StatusCode:                          statusCode,
		RawResponse:                         rawResponse,
		SearchRoles200ApplicationJSONObject: searchRoles200ApplicationJSONObject,
	}

	return out, nil
}