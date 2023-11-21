package provider

import (
	"context"
	"fmt"
	apiclient "terraform-provider-natureremo/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &devicesDataSource{}
	_ datasource.DataSourceWithConfigure = &devicesDataSource{}
)

type devicesDataSource struct {
	client *apiclient.Client
}

func NewDevicesDataSource() datasource.DataSource {
	return &devicesDataSource{}
}

func (d *devicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices"
}

func (d *devicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetch the list of Nature Remo devices.",
		Attributes: map[string]schema.Attribute{
			"devices": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Device identifier attribute.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Device name.",
							Computed:    true,
						},
						"temperature_offset": schema.Float64Attribute{
							Description: "Temperature offset value.",
							Computed:    true,
						},
						"humidity_offset": schema.Int64Attribute{
							Description: "Humidity offset value.",
							Computed:    true,
						},
						"firmware_version": schema.StringAttribute{
							Description: "Firmware version for the device.",
							Computed:    true,
						},
						"mac_address": schema.StringAttribute{
							Description: "MAC address for the device.",
							Computed:    true,
						},
						"bt_mac_address": schema.StringAttribute{
							Description: "Bluetooth MAC address for the device.",
							Computed:    true,
						},
						"serial_number": schema.StringAttribute{
							Description: "Serial number for the device.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *devicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state devicesDataSourceModel

	devices, err := d.client.GetAllDevices(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Nature Remo Devices",
			"Could not read devices, unexpected error: "+err.Error(),
		)
		return
	}

	for _, d := range devices {
		deviceState := deviceDataSourceModel{
			ID:                types.StringValue(d.Id),
			Name:              types.StringValue(d.Name),
			TemperatureOffset: types.Float64Value(d.TemperatureOffset),
			HumidityOffset:    types.Int64Value(d.HumidityOffset),
			FirmwareVersion:   types.StringValue(d.FirmwareVersion),
			MacAddress:        types.StringValue(d.MacAddress),
			BtMacAddress:      types.StringValue(d.BtMacAddress),
			SerialNumber:      types.StringValue(d.SerialNumber),
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

	client, ok := req.ProviderData.(*apiclient.Client)
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
	Devices []deviceDataSourceModel `tfsdk:"devices"`
}

type deviceDataSourceModel struct {
	ID                types.String  `tfsdk:"id"`
	Name              types.String  `tfsdk:"name"`
	TemperatureOffset types.Float64 `tfsdk:"temperature_offset"`
	HumidityOffset    types.Int64   `tfsdk:"humidity_offset"`
	FirmwareVersion   types.String  `tfsdk:"firmware_version"`
	MacAddress        types.String  `tfsdk:"mac_address"`
	BtMacAddress      types.String  `tfsdk:"bt_mac_address"`
	SerialNumber      types.String  `tfsdk:"serial_number"`
}
