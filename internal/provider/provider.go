// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tenntenn/natureremo"
)

// Ensure natureremo satisfies various provider interfaces.
var _ provider.Provider = &natureremoProvider{}

// natureremo defines the provider implementation.
type natureremoProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// natureremoProviderModel describes the provider data model.
type natureremoProviderModel struct {
	AccessToken types.String `tfsdk:"access_token"`
}

func (p *natureremoProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "natureremo"
	resp.Version = p.version
}

func (p *natureremoProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with Nature Remo API.",
		Attributes: map[string]schema.Attribute{
			"access_token": schema.StringAttribute{
				Description: "Access token for Nature Remo. May also be provided via NATURE_REMO_ACCESS_TOKEN environment variable.",
				Required:    true,
				Sensitive:   true,
			},
		},
	}
}

func (p *natureremoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Nature Remo client")
	var data natureremoProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.AccessToken.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("access_token"),
			"Unknown Access Token for Nature Remo API",
			"The provider cannot create the Nature Remo client as there is an unknown access token.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	accessToken := os.Getenv("NATURE_REMO_ACCESS_TOKEN")

	if !data.AccessToken.IsNull() {
		accessToken = data.AccessToken.ValueString()
	}

	if accessToken == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("access_token"),
			"Missing Access Token for Nature Remo API",
			"The provider cannot create the Nature Remo client as there is an missing access token.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "access_token", accessToken)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "access_token")
	tflog.Debug(ctx, "Creating Nature Remo client")

	client := natureremo.NewClient(accessToken)

	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Info(ctx, "Configured Nature Remo client", map[string]any{"success": true})
}

func (p *natureremoProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewUserResource,
		NewDeviceResource,
	}
}

func (p *natureremoProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewUserDataSource,
		NewDevicesDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &natureremoProvider{
			version: version,
		}
	}
}
