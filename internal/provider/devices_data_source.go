package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tenntenn/natureremo"
)

var (
	_ datasource.DataSource              = &devicesDataSource{}
	_ datasource.DataSourceWithConfigure = &devicesDataSource{}
)

type devicesDataSource struct {
	client *natureremo.Client
}

func NewDevicesDataSource() datasource.DataSource {
	return &devicesDataSource{}
}

func (d *devicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices"
}

func (d *devicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"devices": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"temperature_offset": schema.Int64Attribute{
							Computed: true,
						},
						"humidity_offset": schema.Int64Attribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"updated_at": schema.StringAttribute{
							Computed: true,
						},
						"firmware_version": schema.StringAttribute{
							Computed: true,
						},
						"mac_address": schema.StringAttribute{
							Computed: true,
						},
						"bt_mac_address": schema.StringAttribute{
							Computed: true,
						},
						"serial_number": schema.StringAttribute{
							Computed: true,
						},
						"users": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"nickname": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *devicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state devicesDataSourceModel

	devices, err := d.client.DeviceService.GetAll(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Reading Nature Remo Devices",
			"Could not read devices, unexpected error: "+err.Error(),
		)
		return
	}

	for _, d := range devices {
		users := make([]userDataSourceModel, 0, len(d.Users))
		for _, u := range d.Users {
			users = append(users, userDataSourceModel{
				ID:       types.StringValue(u.ID),
				Nickname: types.StringValue(u.Nickname),
			})
		}
		deviceState := deviceModel{
			ID:                types.StringValue(d.ID),
			Name:              types.StringValue(d.Name),
			TemperatureOffset: types.Int64Value(d.TemperatureOffset),
			HumidityOffset:    types.Int64Value(d.HumidityOffset),
			CreatedAt:         types.StringValue(d.CreatedAt.Format(time.RFC3339)),
			UpdatedAt:         types.StringValue(d.UpdatedAt.Format(time.RFC3339)),
			FirmwareVersion:   types.StringValue(d.FirmwareVersion),
			MacAddress:        types.StringValue(d.MacAddress),
			BtMacAddress:      types.StringValue(d.BtMacAddress),
			SerialNumber:      types.StringValue(d.SerialNumber),
			Users:             users,
		}
		state.Devices = append(state.Devices, deviceState)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *devicesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.client = client
}

type devicesDataSourceModel struct {
	Devices []deviceModel `tfsdk:"devices"`
}

type deviceModel struct {
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
