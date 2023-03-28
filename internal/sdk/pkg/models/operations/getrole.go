// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"epilotpermissions/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetRoleRequest struct {
	RoleID string `pathParam:"style=simple,explode=false,name=roleId"`
}

type GetRoleResponse struct {
	ContentType string
	// ok
	Role        *shared.Role
	StatusCode  int
	RawResponse *http.Response
}
