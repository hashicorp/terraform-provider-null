package provider

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/terraform-providers/terraform-provider-null/internal/planmodifiers"
)

var (
	_ resource.Resource = (*nullResource)(nil)
)

func NewNullResource() resource.Resource {
	return &nullResource{}
}

type nullResource struct{}

func (n *nullResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource"
}

func (n *nullResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: `The ` + "`null_resource`" + ` resource implements the standard resource lifecycle but takes no further action.

The ` + "`triggers`" + ` argument allows specifying an arbitrary set of values that, when changed, will cause the resource to be replaced.`,
		Attributes: map[string]tfsdk.Attribute{
			"triggers": {
				Description: "A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.",
				Type:        types.MapType{ElemType: types.StringType},
				Optional:    true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					planmodifiers.RequiresReplaceIfValuesNotNull(),
				},
			},

			"id": {
				Description: "This is set to a random value at create time.",
				Computed:    true,
				Type:        types.StringType,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					resource.UseStateForUnknown(),
				},
			},
		},
	}, nil
}

func (n *nullResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	diags := resp.State.SetAttribute(ctx, path.Root("id"), fmt.Sprintf("%d", rand.Int()))
	resp.Diagnostics.Append(diags...)
}

func (n *nullResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

}

func (n *nullResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var model nullModelV0

	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (n *nullResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}

type nullModelV0 struct {
	Triggers types.Map    `tfsdk:"triggers"`
	ID       types.String `tfsdk:"id"`
}
