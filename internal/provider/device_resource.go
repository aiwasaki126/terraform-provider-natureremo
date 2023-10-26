package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tenntenn/natureremo"
)

var (
	_ resource.Resource                = &deviceResource{}
	_ resource.ResourceWithConfigure   = &deviceResource{}
	_ resource.ResourceWithImportState = &deviceResource{}
)

type deviceResource struct {
	client *natureremo.Client
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
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"temperature_offset": schema.Int64Attribute{
				Required: true,
			},
			"humidity_offset": schema.Int64Attribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"firmware_version": schema.StringAttribute{
				Computed: true,
			},
			"mac_address": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"bt_mac_address": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"serial_number": schema.StringAttribute{
				Computed: true,
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
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"nickname": schema.StringAttribute{
							Computed: true,
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

	devices, err := r.client.DeviceService.GetAll(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Nature Remo Devices",
			"Could not read devices, unexpected error: "+err.Error(),
		)
	}

	var device *natureremo.Device
	for _, d := range devices {
		if d.ID == state.ID.ValueString() {
			device = d
			break
		}
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
			ID:       types.StringValue(u.ID),
			Nickname: types.StringValue(u.Nickname),
		})
	}
	state = deviceResourceModel{
		ID:                types.StringValue(device.ID),
		Name:              types.StringValue(device.Name),
		TemperatureOffset: types.Int64Value(device.TemperatureOffset),
		HumidityOffset:    types.Int64Value(device.HumidityOffset),
		CreatedAt:         types.StringValue(device.CreatedAt.Format(time.RFC3339)),
		UpdatedAt:         types.StringValue(device.UpdatedAt.Format(time.RFC3339)),
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

	requestBody := &natureremo.Device{
		DeviceCore: natureremo.DeviceCore{
			ID:                plan.ID.ValueString(),
			Name:              plan.Name.ValueString(),
			TemperatureOffset: plan.TemperatureOffset.ValueInt64(),
			HumidityOffset:    plan.HumidityOffset.ValueInt64(),
		},
	}

	var device *natureremo.Device
	var err error
	// Update device name
	_, err = r.client.DeviceService.Update(ctx, requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Nature Remo Device Name",
			"Could not update device, unexpexted error: "+err.Error(),
		)
		return
	}
	// Update device temperature offset
	_, err = r.client.DeviceService.UpdateTemperatureOffset(ctx, requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Nature Remo Device Temperature Offset",
			"Could not update device, unexpexted error: "+err.Error(),
		)
		return
	}
	// Update device humidity offset
	device, err = r.client.DeviceService.UpdateHumidityOffset(ctx, requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Nature Remo Device Humidity Offset",
			"Could not update device, unexpexted error: "+err.Error(),
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

	users := make([]userDataSourceModel, 0)
	for _, u := range device.Users {
		users = append(users, userDataSourceModel{
			ID:       types.StringValue(u.ID),
			Nickname: types.StringValue(u.Nickname),
		})
	}
	plan = deviceResourceModel{
		ID:                types.StringValue(device.ID),
		Name:              types.StringValue(device.Name),
		TemperatureOffset: types.Int64Value(device.TemperatureOffset),
		HumidityOffset:    types.Int64Value(device.HumidityOffset),
		CreatedAt:         types.StringValue(device.CreatedAt.Format(time.RFC3339)),
		UpdatedAt:         types.StringValue(device.UpdatedAt.Format(time.RFC3339)),
		FirmwareVersion:   types.StringValue(device.FirmwareVersion),
		MacAddress:        types.StringValue(device.MacAddress),
		BtMacAddress:      types.StringValue(device.BtMacAddress),
		SerialNumber:      types.StringValue(device.SerialNumber),
		Users:             users,
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

	err := r.client.DeviceService.Delete(ctx, &natureremo.Device{
		DeviceCore: natureremo.DeviceCore{
			ID: state.ID.ValueString(),
		}})
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

	client, ok := req.ProviderData.(*natureremo.Client)
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
	TemperatureOffset types.Int64           `tfsdk:"temperature_offset"`
	HumidityOffset    types.Int64           `tfsdk:"humidity_offset"`
	CreatedAt         types.String          `tfsdk:"created_at"`
	UpdatedAt         types.String          `tfsdk:"updated_at"`
	FirmwareVersion   types.String          `tfsdk:"firmware_version"`
	MacAddress        types.String          `tfsdk:"mac_address"`
	BtMacAddress      types.String          `tfsdk:"bt_mac_address"`
	SerialNumber      types.String          `tfsdk:"serial_number"`
	Users             []userDataSourceModel `tfsdk:"users"`
}
