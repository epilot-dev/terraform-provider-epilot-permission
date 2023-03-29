// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"epilot-permission/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutRoleRequest struct {
	Role   *shared.Role `request:"mediaType=application/json"`
	RoleID string       `pathParam:"style=simple,explode=false,name=roleId"`
}

type PutRoleResponse struct {
	ContentType string
	// ok
	Role        *shared.Role
	StatusCode  int
	RawResponse *http.Response
}
