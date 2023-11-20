package provider

import (
	"context"
	"fmt"
	apiclient "terraform-provider-natureremo/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &deviceResource{}
	_ resource.ResourceWithConfigure   = &deviceResource{}
	_ resource.ResourceWithImportState = &deviceResource{}
)

type deviceResource struct {
	client *apiclient.Client
}

func NewDeviceResource() resource.Resource {
	return &deviceResource{}
}

func (r *deviceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device"
}

func (r *deviceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Device identifier attribute.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Device name.",
				Required:    true,
			},
			"temperature_offset": schema.Float64Attribute{
				Description: "Temperature offset value.",
				Required:    true,
			},
			"humidity_offset": schema.Int64Attribute{
				Description: "Humidity offset value.",
				Required:    true,
			},
			"firmware_version": schema.StringAttribute{
				Description: "Firmware version for the device.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"mac_address": schema.StringAttribute{
				Description: "MAC address for the device.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"bt_mac_address": schema.StringAttribute{
				Description: "Bluetooth MAC address for the device.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"serial_number": schema.StringAttribute{
				Description: "Serial number for the device.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"users": schema.ListNestedAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
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
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *deviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
}

// Read refreshes the Terraform state with the latest data.
func (r *deviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state deviceResourceModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device, err := r.client.GetDevice(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Nature Remo Devices",
			"Could not read devices, unexpected error: "+err.Error(),
		)
	}
	if device == nil {
		resp.Diagnostics.AddError(
			"Error Reading Nature Remo Device",
			"Could not find the specified device",
		)
		return
	}

	users := make([]userDataSourceModel, 0, len(device.Users))
	for _, u := range device.Users {
		users = append(users, userDataSourceModel{
			ID:       types.StringValue(u.Id),
			Nickname: types.StringValue(u.Nickname),
		})
	}
	state = deviceResourceModel{
		ID:                types.StringValue(device.Id),
		Name:              types.StringValue(device.Name),
		TemperatureOffset: types.Float64Value(device.TemperatureOffset),
		HumidityOffset:    types.Int64Value(device.HumidityOffset),
		FirmwareVersion:   types.StringValue(device.FirmwareVersion),
		MacAddress:        types.StringValue(device.MacAddress),
		BtMacAddress:      types.StringValue(device.BtMacAddress),
		SerialNumber:      types.StringValue(device.SerialNumber),
		Users:             users,
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *deviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan deviceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update device name
	device, err := r.client.UpdateDevice(ctx, plan.ID.ValueString(), plan.Name.ValueString(), plan.HumidityOffset.ValueInt64(), plan.TemperatureOffset.ValueFloat64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Nature Remo Device Name",
			"Could not update device, unexpected error: "+err.Error(),
		)
		return
	}

	if device == nil {
		resp.Diagnostics.AddError(
			"Error Reading Nature Remo Device",
			"Could not find the specified device",
		)
		return
	}

	plan = deviceResourceModel{
		ID:                types.StringValue(device.Id),
		Name:              types.StringValue(device.Name),
		TemperatureOffset: types.Float64Value(device.TemperatureOffset),
		HumidityOffset:    types.Int64Value(device.HumidityOffset),
		FirmwareVersion:   plan.FirmwareVersion,
		MacAddress:        plan.MacAddress,
		BtMacAddress:      plan.BtMacAddress,
		SerialNumber:      plan.SerialNumber,
		Users:             plan.Users,
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *deviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state deviceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteDevice(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Nature Remo Device",
			"Could not delete device, unexpected error: "+err.Error(),
		)
		return
	}
}

// Configure adds the provider configured client to the resource.
func (r *deviceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *deviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

type deviceResourceModel struct {
	ID                types.String          `tfsdk:"id"`
	Name              types.String          `tfsdk:"name"`
	TemperatureOffset types.Float64         `tfsdk:"temperature_offset"`
	HumidityOffset    types.Int64           `tfsdk:"humidity_offset"`
	FirmwareVersion   types.String          `tfsdk:"firmware_version"`
	MacAddress        types.String          `tfsdk:"mac_address"`
	BtMacAddress      types.String          `tfsdk:"bt_mac_address"`
	SerialNumber      types.String          `tfsdk:"serial_number"`
	Users             []userDataSourceModel `tfsdk:"users"`
}
