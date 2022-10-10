package provider

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource = (*nullDataSource)(nil)
)

func NewNullDataSource() datasource.DataSource {
	return &nullDataSource{}
}

type nullDataSource struct{}

func (n nullDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_data_source"
}

func (n nullDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		DeprecationMessage: "The null_data_source was historically used to construct intermediate values to re-use elsewhere " +
			"in configuration, the same can now be achieved using locals",
		Description: `The ` + "`null_data_source`" + ` data source implements the standard data source lifecycle but does not
interact with any external APIs.

Historically, the ` + "`null_data_source`" + ` was typically used to construct intermediate values to re-use elsewhere in configuration. The
same can now be achieved using [locals](https://www.terraform.io/docs/language/values/locals.html).
`,
		Attributes: map[string]tfsdk.Attribute{
			"inputs": {
				Description: "A map of arbitrary strings that is copied into the `outputs` attribute, and accessible directly for interpolation.",
				Type:        types.MapType{ElemType: types.StringType},
				Optional:    true,
			},
			"outputs": {
				Description: "After the data source is \"read\", a copy of the `inputs` map.",
				Type:        types.MapType{ElemType: types.StringType},
				Computed:    true,
			},
			"random": {
				Description: "A random value. This is primarily for testing and has little practical use; prefer the [hashicorp/random provider](https://registry.terraform.io/providers/hashicorp/random) for more practical random number use-cases.",
				Type:        types.StringType,
				Computed:    true,
			},
			"has_computed_default": {
				Description: "If set, its literal value will be stored and returned. If not, its value defaults to `\"default\"`. This argument exists primarily for testing and has little practical use.",
				Type:        types.StringType,
				Optional:    true,
				Computed:    true,
			},

			"id": {
				Description:        "This attribute is only present for some legacy compatibility issues and should not be used. It will be removed in a future version.",
				Computed:           true,
				DeprecationMessage: "This attribute is only present for some legacy compatibility issues and should not be used. It will be removed in a future version.",
				Type:               types.StringType,
			},
		},
	}, nil
}

func (n nullDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config nullDataSourceModelV0

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	config.Outputs = config.Inputs
	config.Random = types.String{Value: fmt.Sprintf("%d", rand.Int())}

	if config.HasComputedDefault.IsNull() {
		config.HasComputedDefault.Null = false
		config.HasComputedDefault = types.String{Value: "default"}
	}
	config.ID = types.String{Value: "static"}
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
