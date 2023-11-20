package provider

import (
	"context"
	"fmt"
	apiclient "terraform-provider-natureremo/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var (
	_ resource.Resource                = &userResource{}
	_ resource.ResourceWithConfigure   = &userResource{}
	_ resource.ResourceWithImportState = &userResource{}
)

type userResourceModel struct {
	ID           types.String `tfsdk:"id"`
	Nickname     types.String `tfsdk:"nickname"`
	Country      types.String `tfsdk:"country"`
	DistanceUnit types.String `tfsdk:"distance_unit"`
	TempUnit     types.String `tfsdk:"temp_unit"`
}

type userResource struct {
	client *apiclient.Client
}

func NewUserResource() resource.Resource {
	return &userResource{}
}

func (r *userResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (r *userResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages an user.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier of user.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"nickname": schema.StringAttribute{
				Description: "Nickname of user.",
				Required:    true,
			},
			"country": schema.StringAttribute{
				Description: "Country",
				Optional:    true,
			},
			"distance_unit": schema.StringAttribute{
				Description: "Distance unit",
				Optional:    true,
			},
			"temp_unit": schema.StringAttribute{
				Description: "Temperature unit",
				Optional:    true,
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *userResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
}

// Read refreshes the Terraform state with the latest data.
func (r *userResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state userResourceModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	user, err := r.client.GetProfile(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Nature Remo User",
			"Could not read user, unexpected error: "+err.Error(),
		)
	}

	state.ID = types.StringValue(user.Id)
	state.Nickname = types.StringValue(user.Nickname)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *userResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan userResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	user, err := r.client.UpdateProfile(ctx,
		plan.ID.ValueString(),
		plan.Nickname.ValueString(),
		plan.Country.ValueString(),
		plan.DistanceUnit.ValueString(),
		plan.TempUnit.ValueString(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Nature Remo User",
			"Could not update user, unexpected error"+err.Error(),
		)
	}

	plan = userResourceModel{
		ID:           types.StringValue(user.Id),
		Nickname:     types.StringValue(user.Nickname),
		Country:      types.StringValue(plan.Country.ValueString()),
		DistanceUnit: types.StringValue(plan.DistanceUnit.ValueString()),
		TempUnit:     types.StringValue(plan.TempUnit.ValueString()),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *userResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

// Configure adds the provider configured client to the resource.
func (r *userResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*apiclient.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *natureremo.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *userResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
