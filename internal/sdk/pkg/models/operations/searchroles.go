// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"epilotpermissions/internal/sdk/pkg/models/shared"
	"net/http"
)

type SearchRolesRequest struct {
	Request *shared.RoleSearchInput `request:"mediaType=application/json"`
}

// SearchRoles200ApplicationJSON - ok
type SearchRoles200ApplicationJSON struct {
	Hits    *float64      `json:"hits,omitempty"`
	Results []shared.Role `json:"results,omitempty"`
}

type SearchRolesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// ok
	SearchRoles200ApplicationJSONObject *SearchRoles200ApplicationJSON
}
