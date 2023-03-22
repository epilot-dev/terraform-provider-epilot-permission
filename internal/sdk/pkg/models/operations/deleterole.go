// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"epilotpermissions/internal/sdk/pkg/models/shared"
	"fmt"
	"net/http"
)

type DeleteRolePathParams struct {
	RoleID string `pathParam:"style=simple,explode=false,name=roleId"`
}

func NewDeleteRolePathParams(input interface{}) (*DeleteRolePathParams, error) {
	mapInput, ok := input.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("deleteRolePathParams: Expected input to be a map[string]interface{}")
	}
	if _, ok = mapInput["roleId"]; !ok {
		return nil, fmt.Errorf("deleteRolePathParams: RoleID is required, but was not found")
	}
	var roleID string
	roleID, ok = mapInput["roleId"].(string)
	if !ok {
		return nil, fmt.Errorf("deleteRolePathParams: unexpected type for RoleID. Expected string but was %T", mapInput["roleId"])
	}
	out := &DeleteRolePathParams{
		RoleID: roleID,
	}

	return out, nil
}

type DeleteRoleRequest struct {
	PathParams DeleteRolePathParams
}

func NewDeleteRoleRequest(input interface{}) (*DeleteRoleRequest, error) {
	mapInput, ok := input.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("deleteRoleRequest: Expected input to be a map[string]interface{}")
	}
	if _, ok = mapInput["PathParams"]; !ok {
		return nil, fmt.Errorf("deleteRoleRequest: PathParams is required, but was not found")
	}
	pathParamsPtr, err := NewDeleteRolePathParams(mapInput["PathParams"])
	if err != nil {
		return nil, err
	}
	pathParams := *pathParamsPtr
	out := &DeleteRoleRequest{
		PathParams: pathParams,
	}

	return out, nil
}

type DeleteRoleResponse struct {
	ContentType string
	// ok
	Role        interface{}
	StatusCode  int
	RawResponse *http.Response
}

func NewDeleteRoleResponse(input interface{}) (*DeleteRoleResponse, error) {
	mapInput, ok := input.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("deleteRoleResponse: Expected input to be a map[string]interface{}")
	}
	if _, ok = mapInput["ContentType"]; !ok {
		return nil, fmt.Errorf("deleteRoleResponse: ContentType is required, but was not found")
	}
	var contentType string
	contentType, ok = mapInput["ContentType"].(string)
	if !ok {
		return nil, fmt.Errorf("deleteRoleResponse: unexpected type for ContentType. Expected string but was %T", mapInput["ContentType"])
	}
	var role interface{}
	pickFirstValidRole := make([]func(interface{}) (interface{}, error), 0)
	pickFirstValidRole = append(pickFirstValidRole, func(sharedUserRoleInput interface{}) (interface{}, error) {
		var sharedUserRoleOutput *shared.UserRole
		if sharedUserRoleInput != nil {
			roleTmp, err := shared.NewUserRole(sharedUserRoleInput)
			if err != nil {
				return nil, err
			}
			sharedUserRoleOutput = roleTmp
		}
		return sharedUserRoleOutput, nil
	})
	pickFirstValidRole = append(pickFirstValidRole, func(sharedOrgRoleInput interface{}) (interface{}, error) {
		var sharedOrgRoleOutput *shared.OrgRole
		if sharedOrgRoleInput != nil {
			roleTmp1, err := shared.NewOrgRole(sharedOrgRoleInput)
			if err != nil {
				return nil, err
			}
			sharedOrgRoleOutput = roleTmp1
		}
		return sharedOrgRoleOutput, nil
	})
	pickFirstValidRole = append(pickFirstValidRole, func(sharedShareRoleInput interface{}) (interface{}, error) {
		var sharedShareRoleOutput *shared.ShareRole
		if sharedShareRoleInput != nil {
			roleTmp2, err := shared.NewShareRole(sharedShareRoleInput)
			if err != nil {
				return nil, err
			}
			sharedShareRoleOutput = roleTmp2
		}
		return sharedShareRoleOutput, nil
	})
	pickFirstValidRole = append(pickFirstValidRole, func(sharedPartnerRoleInput interface{}) (interface{}, error) {
		var sharedPartnerRoleOutput *shared.PartnerRole
		if sharedPartnerRoleInput != nil {
			roleTmp3, err := shared.NewPartnerRole(sharedPartnerRoleInput)
			if err != nil {
				return nil, err
			}
			sharedPartnerRoleOutput = roleTmp3
		}
		return sharedPartnerRoleOutput, nil
	})
	for _, roleChecker := range pickFirstValidRole {
		var roleErr error
		role, roleErr = roleChecker(mapInput["Role"])
		if roleErr == nil {
			break
		}
	}
	if role == nil {
		return nil, fmt.Errorf("deleteRoleResponse: unexpected type for Role. Expected one of shared.UserRole, shared.OrgRole, shared.ShareRole, shared.PartnerRole to be valid. Got %#v", mapInput["Role"])
	}
	if _, ok = mapInput["StatusCode"]; !ok {
		return nil, fmt.Errorf("deleteRoleResponse: StatusCode is required, but was not found")
	}
	if _, ok = mapInput["StatusCode"]; !ok {
		return nil, fmt.Errorf("deleteRoleResponse: StatusCode is required, but was not found")
	}
	var statusCodeFloat64 float64
	statusCodeFloat64, ok = mapInput["StatusCode"].(float64)
	if !ok {
		return nil, fmt.Errorf("deleteRoleResponse: unexpected type for StatusCode. Expected float64 but was %T", mapInput["StatusCode"])
	}
	var statusCode int
	if statusCodeFloat64 != float64(int(statusCodeFloat64)) {
		return nil, fmt.Errorf("deleteRoleResponse: unexpected value for integer StatusCode. Got %#v", mapInput["StatusCode"])
	} else {
		statusCode = int(statusCodeFloat64)
	}
	rawResponse := new(http.Response)
	if _, ok = mapInput["RawResponse"]; !ok {
		rawResponse = nil
	} else {
		*rawResponse, ok = mapInput["RawResponse"].(http.Response)
		if !ok {
			return nil, fmt.Errorf("deleteRoleResponse: unexpected type for RawResponse. Expected http.Response but was %T", mapInput["RawResponse"])
		}
	}
	out := &DeleteRoleResponse{
		ContentType: contentType,
		Role:        role,
		StatusCode:  statusCode,
		RawResponse: rawResponse,
	}

	return out, nil
}
