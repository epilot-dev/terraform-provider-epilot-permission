// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type Grant struct {
	Action   types.String `tfsdk:"action"`
	Effect   types.String `tfsdk:"effect"`
	Resource types.String `tfsdk:"resource"`
}
