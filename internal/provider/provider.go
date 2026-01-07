// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = (*nullProvider)(nil)

func New() provider.Provider {
	return &nullProvider{}
}

type nullProvider struct{}

func (p *nullProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "null"
}

func (p *nullProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

func (p *nullProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewNullDataSource,
	}
}

func (p *nullProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewNullResource,
	}
}

func (p *nullProvider) Schema(context.Context, provider.SchemaRequest, *provider.SchemaResponse) {
}
