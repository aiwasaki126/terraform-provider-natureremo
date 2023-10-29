package provider

// TODO: It's better to split  resource, data_source into 3 applicant kind, aircon, tv, light.

// var (
// 	_ datasource.DataSource              = &appliancesDataSource{}
// 	_ datasource.DataSourceWithConfigure = &appliancesDataSource{}
// )

// type appliancesDataSource struct {
// 	client *natureremo.Client
// }

// func NewAppliancesDataSource() datasource.DataSource {
// 	return &appliancesDataSource{}
// }

// func (d *appliancesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
// 	resp.TypeName = req.ProviderTypeName + "_appliances"
// }

// func (d *appliancesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
// 	resp.Schema = schema.Schema{
// 		Attributes: map[string]schema.Attribute{
// 			"appliances": schema.ListNestedAttribute{
// 				Computed: true,
// 				NestedObject: schema.NestedAttributeObject{
// 					Attributes: map[string]schema.Attribute{
// 						"id": schema.StringAttribute{
// 							Computed: true,
// 						},
// 						"type": schema.StringAttribute{
// 							Computed: true,
// 						},
// 						"device": schema.NestedAttributeObject{
// 							Computed:   true,
// 							Attributes: deviceDataSourceAttribute,
// 						},

// 						"model": schema.ObjectAttribute{
// 							Computed: true,
// 						},
// 						"nickname": schema.StringAttribute{
// 							Computed: true,
// 						},
// 						"image": schema.StringAttribute{
// 							Computed: true,
// 						},
// 						"signals": schema.ListNestedAttribute{
// 							Computed: true,
// 						},
// 						"settings": schema.NestedAttributeObject{
// 							Computed:   true,
// 							Attributes: map[string]schema.Attribute{},
// 						},
// 						"aircon": schema.ObjectAttribute{
// 							Computed: true,
// 						},
// 						"tv": schema.ObjectAttribute{
// 							Computed: true,
// 						},
// 						"light": schema.ObjectAttribute{
// 							Computed: true,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// func (d *appliancesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
// 	var state appliancesDataSourceModel

// 	appliances, err := d.client.ApplianceService.GetAll(ctx)
// 	if err != nil {
// 		resp.Diagnostics.AddError(
// 			"Error Reading Nature Remo Appliances",
// 			"Could not read appliances, unexpected error: "+err.Error(),
// 		)
// 		return
// 	}
// 	for _, d := range appliances {
// 		applianceState := applianceDataSourceModel{
// 			ID: types.StringValue(d.ID),
// 		}
// 		state.Appliances = append(state.Appliances, applianceState)
// 	}

// 	diags := resp.State.Set(ctx, &state)
// 	resp.Diagnostics.Append(diags...)
// 	if resp.Diagnostics.HasError() {
// 		return
// 	}
// }

// func (d *appliancesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
// 	if req.ProviderData == nil {
// 		return
// 	}

// 	client, ok := req.ProviderData.(*natureremo.Client)
// 	if !ok {
// 		resp.Diagnostics.AddError(
// 			"Unexpected Data Source Configure Type",
// 			fmt.Sprintf("Expected *natureremo.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
// 		)

// 		return
// 	}

// 	d.client = client
// }

// type appliancesDataSourceModel struct {
// 	Appliances []applianceDataSourceModel `tfsdk:"appliances"`
// }

// type applianceDataSourceModel struct {
// 	ID             types.String    `tfsdk:"id"`
// 	Type           types.String    `tfsdk:"type"`
// 	Device         *deviceModel    `tfsdk:"device"`
// 	Model          *model          `tfsdk:"model"`
// 	Nickname       types.String    `tfsdk:"nickname"`
// 	Image          types.String    `tfsdk:"image"`
// 	Signals        []*signal       `tfsdk:"signals"`
// 	AirConSettings *airConSettings `tfsdk:"settings"`
// 	AirCon         *airCon         `tfsdk:"aircon"`
// 	TV             *tv             `tfsdk:"tv"`
// 	Light          *light          `tfsdk:"light"`
// }

// type model struct {
// 	ID           types.String `tfsdk:"id"`
// 	Country      types.String `tfsdk:"country"`
// 	Manufacturer types.String `tfsdk:"manufacturer"`
// 	RemoteName   types.String `tfsdk:"remote_name"`
// 	Name         types.String `tfsdk:"name"`
// 	Image        types.String `tfsdk:"image"`
// }

// type signal struct {
// 	ID    string `json:"id"`
// 	Name  string `json:"name"`
// 	Image string `json:"image"`
// }

// type airCon struct {
// 	Range           *AirConRange    `json:"range"`
// 	TemperatureUnit TemperatureUnit `json:"temperautre_unit"`
// }

// type airConRange struct {
// 	Modes        map[OperationMode]*AirConRangeMode `json:"modes"`
// 	FixedButtons []Button                           `json:"fixedButtons"`
// }
