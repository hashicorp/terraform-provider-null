// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource = (*nullDataSource)(nil)
)

func NewNullDataSource() datasource.DataSource {
	return &nullDataSource{}
}

type nullDataSource struct{}

func (n *nullDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_data_source"
}

func (n *nullDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		DeprecationMessage: "The null_data_source was historically used to construct intermediate values to re-use elsewhere " +
			"in configuration, the same can now be achieved using locals or the terraform_data resource type in Terraform 1.4 and later.",
		Description: `The ` + "`null_data_source`" + ` data source implements the standard data source lifecycle but does not
interact with any external APIs.

Historically, the ` + "`null_data_source`" + ` was typically used to construct intermediate values to re-use elsewhere in configuration. The
same can now be achieved using [locals](https://developer.hashicorp.com/terraform/language/values/locals) ` +
			`or the [terraform_data resource type](https://developer.hashicorp.com/terraform/language/resources/terraform-data) in Terraform 1.4 and later.`,
		Attributes: map[string]schema.Attribute{
			"inputs": schema.MapAttribute{
				Description: "A map of arbitrary strings that is copied into the `outputs` attribute, and accessible directly for interpolation.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"outputs": schema.MapAttribute{
				Description: "After the data source is \"read\", a copy of the `inputs` map.",
				ElementType: types.StringType,
				Computed:    true,
			},
			"random": schema.StringAttribute{
				Description: "A random value. This is primarily for testing and has little practical use; prefer the [hashicorp/random provider](https://registry.terraform.io/providers/hashicorp/random) for more practical random number use-cases.",
				Computed:    true,
			},
			"has_computed_default": schema.StringAttribute{
				Description: "If set, its literal value will be stored and returned. If not, its value defaults to `\"default\"`. This argument exists primarily for testing and has little practical use.",
				Optional:    true,
				Computed:    true,
			},

			"id": schema.StringAttribute{
				Description:        "This attribute is only present for some legacy compatibility issues and should not be used. It will be removed in a future version.",
				Computed:           true,
				DeprecationMessage: "This attribute is only present for some legacy compatibility issues and should not be used. It will be removed in a future version.",
			},
		},
	}
}

func (n *nullDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config nullDataSourceModelV0

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	config.Outputs = config.Inputs
	config.Random = types.StringValue(fmt.Sprintf("%d", rand.Int()))

	if config.HasComputedDefault.IsNull() {
		config.HasComputedDefault = types.StringValue("default")
	}
	config.ID = types.StringValue("static")
	diags = resp.State.Set(ctx, config)
	resp.Diagnostics.Append(diags...)
}

type nullDataSourceModelV0 struct {
	Inputs             types.Map    `tfsdk:"inputs"`
	Outputs            types.Map    `tfsdk:"outputs"`
	Random             types.String `tfsdk:"random"`
	HasComputedDefault types.String `tfsdk:"has_computed_default"`
	ID                 types.String `tfsdk:"id"`
}
