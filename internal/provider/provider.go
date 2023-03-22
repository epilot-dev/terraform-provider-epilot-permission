// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"epilotpermissions/internal/sdk"
	"epilotpermissions/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &EpilotpermissionsProvider{}

type EpilotpermissionsProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// EpilotpermissionsProviderModel describes the provider data model.
type EpilotpermissionsProviderModel struct {
	ServerURL     types.String `tfsdk:"server_url"`
	Authorization types.String `tfsdk:"authorization"`
	APIKey        types.String `tfsdk:"apikey"`
}

func (p *EpilotpermissionsProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "epilotpermissions"
	resp.Version = p.version
}

func (p *EpilotpermissionsProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to https://permissions.sls.epilot.io)",
				Optional:            true,
			},
			"authorization": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"apikey": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *EpilotpermissionsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data EpilotpermissionsProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "https://permissions.sls.epilot.io"
	}

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(shared.Security{
			EpilotAuth: &shared.SchemeEpilotAuth{
				Authorization: data.Authorization.ValueString(),
			},
		}),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *EpilotpermissionsProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAssignmentResource,
		NewRoleResource,
	}
}

func (p *EpilotpermissionsProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &EpilotpermissionsProvider{
			version: version,
		}
	}
}
