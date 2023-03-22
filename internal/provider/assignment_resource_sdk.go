// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"epilotpermissions/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *AssignmentResourceModel) ToSDKType() interface{} {
	// Not Used

	return nil

}

func (r *AssignmentResourceModel) RefreshFromSDKType(resp *shared.Assignment) {
	r.Roles = nil
	for _, v := range resp.Roles {
		r.Roles = append(r.Roles, types.StringValue(v))
	}
	if resp.UserID != nil {
		r.UserID = types.StringValue(*resp.UserID)
	} else {
		r.UserID = types.StringNull()
	}

}
