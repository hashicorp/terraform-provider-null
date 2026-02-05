// Copyright IBM Corp. 2017, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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

func (n *nullResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The `null_resource` resource implements the standard resource lifecycle but takes no further action. " +
			"On Terraform 1.4 and later, use the [`terraform_data` resource type](https://developer.hashicorp.com/terraform/language/resources/terraform-data) instead. " +
			"Terraform 1.9 and later support the `moved` configuration block from `null_resource` to `terraform_data`.\n\n" +
			"The `triggers` argument allows specifying an arbitrary set of values that, when changed, will cause the resource to be replaced.",
		Attributes: map[string]schema.Attribute{
			"triggers": schema.MapAttribute{
				Description: "A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.",
				ElementType: types.StringType,
				Optional:    true,
				PlanModifiers: []planmodifier.Map{
					planmodifiers.RequiresReplaceIfValuesNotNull(),
				},
			},

			"id": schema.StringAttribute{
				Description: "This is set to a random value at create time.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (n *nullResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var model nullModelV0

	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)

	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)

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
