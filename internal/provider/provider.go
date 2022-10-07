package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func New() provider.Provider {
	return &nullProvider{}
}

var (
	_ provider.Provider             = (*nullProvider)(nil)
	_ provider.ProviderWithMetadata = (*nullProvider)(nil)
)

type nullProvider struct{}

func (p *nullProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "null"
}

func (p *nullProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

func (p *nullProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}

func (p *nullProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *nullProvider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{}, nil
}
