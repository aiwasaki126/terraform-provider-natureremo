// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package gen

import (
	"time"
)

const (
	Oauth2Scopes = "oauth2.Scopes"
)

// Defines values for UpdateProfileParamCountry.
const (
	AU     UpdateProfileParamCountry = "AU"
	CA     UpdateProfileParamCountry = "CA"
	JP     UpdateProfileParamCountry = "JP"
	NZ     UpdateProfileParamCountry = "NZ"
	OTHERS UpdateProfileParamCountry = "OTHERS"
	SG     UpdateProfileParamCountry = "SG"
	US     UpdateProfileParamCountry = "US"
)

// Defines values for UpdateProfileParamDistanceUnit.
const (
	Imperial UpdateProfileParamDistanceUnit = "imperial"
	Metric   UpdateProfileParamDistanceUnit = "metric"
)

// Defines values for UpdateProfileParamTempUnit.
const (
	C UpdateProfileParamTempUnit = "c"
	F UpdateProfileParamTempUnit = "f"
)

// AirConParams defines model for AirConParams_.
type AirConParams struct {
	// AirDirection AC air direction. Empty means automatic.
	AirDirection *string `json:"air_direction,omitempty"`

	// AirDirectionH AC horizontal air direction.
	AirDirectionH *string `json:"air_direction_h,omitempty"`

	// AirVolume AC air volume. Empty means automatic. Numbers express the amount of volume. The range of AirVolumes which the air conditioner accepts depends on the air conditioner model and operation mode. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model and operation mode.
	AirVolume *string `json:"air_volume,omitempty"`

	// Button Button. Specify 'power-off' always if you want the air conditioner powered off. Empty means powered on.
	Button *string `json:"button,omitempty"`

	// OperationMode AC operation mode. The range of operation modes which the air conditioner accepts depends on the air conditioner model. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model.
	OperationMode *string `json:"operation_mode,omitempty"`

	// Temperature Temperature. The temperature in string format. The unit is described in Aircon object. The range of Temperatures which the air conditioner accepts depends on the air conditioner model and operation mode. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model and operation mode.
	Temperature *string `json:"temperature,omitempty"`

	// TemperatureUnit Temperature unit. 'c' or 'f' or '' for unknown.
	TemperatureUnit *string `json:"temperature_unit,omitempty"`
}

// AirconSettingsResponse defines model for AirconSettingsResponse.
type AirconSettingsResponse struct {
	// Button Button. Specify 'power-off' always if you want the air conditioner powered off. Empty means powered on.
	Button *string `json:"button,omitempty"`

	// Dir AC air direction. Empty means automatic.
	Dir *string `json:"dir,omitempty"`

	// Dirh AC horizontal air direction.
	Dirh *string `json:"dirh,omitempty"`

	// Mode AC operation mode. The range of operation modes which the air conditioner accepts depends on the air conditioner model. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model.
	Mode *string `json:"mode,omitempty"`

	// Temp Temperature. The temperature in string format. The unit is described in Aircon object. The range of Temperatures which the air conditioner accepts depends on the air conditioner model and operation mode. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model and operation mode.
	Temp *string `json:"temp,omitempty"`

	// TempUnit Temperature unit. 'c' or 'f' or '' for unknown.
	TempUnit  *string    `json:"temp_unit,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// Vol AC air volume. Empty means automatic. Numbers express the amount of volume. The range of AirVolumes which the air conditioner accepts depends on the air conditioner model and operation mode. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model and operation mode.
	Vol *string `json:"vol,omitempty"`
}

// ApplianceModelAndParam defines model for ApplianceModelAndParam.
type ApplianceModelAndParam struct {
	Model *struct {
		Country      *string `json:"country,omitempty"`
		Id           *string `json:"id,omitempty"`
		Image        *string `json:"image,omitempty"`
		Manufacturer *string `json:"manufacturer,omitempty"`
		Name         *string `json:"name,omitempty"`
		RemoteName   *string `json:"remote_name,omitempty"`
		Series       *string `json:"series,omitempty"`
	} `json:"model,omitempty"`
	Params *struct {
		// Button Button. Specify 'power-off' always if you want the air conditioner powered off. Empty means powered on.
		Button *string `json:"button,omitempty"`

		// Dir AC air direction. Empty means automatic.
		Dir *string `json:"dir,omitempty"`

		// Dirh AC horizontal air direction.
		Dirh *string `json:"dirh,omitempty"`

		// Mode AC operation mode. The range of operation modes which the air conditioner accepts depends on the air conditioner model. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model.
		Mode *string `json:"mode,omitempty"`

		// Temp Temperature. The temperature in string format. The unit is described in Aircon object. The range of Temperatures which the air conditioner accepts depends on the air conditioner model and operation mode. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model and operation mode.
		Temp *string `json:"temp,omitempty"`

		// TempUnit Temperature unit. 'c' or 'f' or '' for unknown.
		TempUnit *string `json:"temp_unit,omitempty"`

		// Vol AC air volume. Empty means automatic. Numbers express the amount of volume. The range of AirVolumes which the air conditioner accepts depends on the air conditioner model and operation mode. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model and operation mode.
		Vol *string `json:"vol,omitempty"`
	} `json:"params,omitempty"`
}

// ApplianceModelAndParams defines model for ApplianceModelAndParams.
type ApplianceModelAndParams = []ApplianceModelAndParam

// ApplianceParams defines model for ApplianceParams.
type ApplianceParams struct {
	// Image Basename of the image file included in the app. Ex: 'ico_ac_1'.
	Image *string `json:"image,omitempty"`

	// Nickname Appliance name.
	Nickname *string `json:"nickname,omitempty"`
}

// ApplianceResponse defines model for ApplianceResponse.
type ApplianceResponse struct {
	AcSmartMode *struct {
		Adjusting *bool `json:"adjusting,omitempty"`
		Area      *int  `json:"area,omitempty"`
		Enabled   *bool `json:"enabled,omitempty"`
	} `json:"ac_smart_mode,omitempty"`
	Aircon *struct {
		Range *struct {
			FixedButtons *[]string `json:"fixedButtons,omitempty"`
			Modes        *map[string]struct {
				Dir  *[]string `json:"dir,omitempty"`
				Dirh *[]string `json:"dirh,omitempty"`
				Temp *[]string `json:"temp,omitempty"`
				Vol  *[]string `json:"vol,omitempty"`
			} `json:"modes,omitempty"`
		} `json:"range,omitempty"`
		TempUnit *string `json:"tempUnit,omitempty"`
	} `json:"aircon,omitempty"`
	AirconSmartEcoMode *struct {
		Adjusting *bool `json:"adjusting,omitempty"`
		Area      *int  `json:"area,omitempty"`
		Enabled   *bool `json:"enabled,omitempty"`
	} `json:"aircon_smart_eco_mode,omitempty"`
	Device *struct {
		BtMacAddress      *string    `json:"bt_mac_address,omitempty"`
		CreatedAt         *time.Time `json:"created_at,omitempty"`
		FirmwareVersion   *string    `json:"firmware_version,omitempty"`
		HumidityOffset    *float32   `json:"humidity_offset,omitempty"`
		Id                *string    `json:"id,omitempty"`
		MacAddress        *string    `json:"mac_address,omitempty"`
		Name              *string    `json:"name,omitempty"`
		SerialNumber      *string    `json:"serial_number,omitempty"`
		TemperatureOffset *float32   `json:"temperature_offset,omitempty"`
		UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	} `json:"device,omitempty"`

	// Echonetlite The ECHONET lite properties fetched from the appliance. See 'Detailed Requirements for ECHONET Device Objects' for more details. ref. https://echonet.jp/spec_object_rl_en/
	Echonetlite *struct {
		GetProperties *map[string]struct {
			Enum *[]string `json:"enum,omitempty"`
			Type *string   `json:"type,omitempty"`
		} `json:"get_properties,omitempty"`
		Identifier *string `json:"identifier,omitempty"`
		Instance   *string `json:"instance,omitempty"`
		Ip         *string `json:"ip,omitempty"`
		Localize   *struct {
			Properties *map[string]struct {
				Enum *map[string]struct {
					Label *string `json:"label,omitempty"`
				} `json:"enum,omitempty"`
			} `json:"properties,omitempty"`
		} `json:"localize,omitempty"`
		RouteType     *string `json:"route_type,omitempty"`
		SetProperties *map[string]struct {
			Enum *[]string `json:"enum,omitempty"`
			Type *string   `json:"type,omitempty"`
		} `json:"set_properties,omitempty"`
		State   *interface{} `json:"state,omitempty"`
		Version *string      `json:"version,omitempty"`
	} `json:"echonetlite,omitempty"`
	Id    *string `json:"id,omitempty"`
	Image *string `json:"image,omitempty"`
	Light *struct {
		Buttons *[]struct {
			Image *string `json:"image,omitempty"`
			Label *string `json:"label,omitempty"`
			Name  *string `json:"name,omitempty"`
		} `json:"buttons,omitempty"`
		State *struct {
			Brightness *string `json:"brightness,omitempty"`
			LastButton *string `json:"last_button,omitempty"`
			Power      *string `json:"power,omitempty"`
		} `json:"state,omitempty"`
	} `json:"light,omitempty"`
	LightProjector *struct {
		Layout *struct {
			Image     *string `json:"image,omitempty"`
			Label     *string `json:"label,omitempty"`
			Name      *string `json:"name,omitempty"`
			Templates *[]struct {
				Image     *string     `json:"image,omitempty"`
				Label     *string     `json:"label,omitempty"`
				Name      *string     `json:"name,omitempty"`
				Templates *[]Template `json:"templates,omitempty"`
				Text      *string     `json:"text,omitempty"`
				Type      *string     `json:"type,omitempty"`
				Uuid      *string     `json:"uuid,omitempty"`
				XSize     *int        `json:"x_size,omitempty"`
				YSize     *int        `json:"y_size,omitempty"`
			} `json:"templates,omitempty"`
			Text  *string `json:"text,omitempty"`
			Type  *string `json:"type,omitempty"`
			Uuid  *string `json:"uuid,omitempty"`
			XSize *int    `json:"x_size,omitempty"`
			YSize *int    `json:"y_size,omitempty"`
		} `json:"layout,omitempty"`
	} `json:"light_projector,omitempty"`
	Model      *interface{} `json:"model,omitempty"`
	MorninPlus *struct {
		Devices *[]struct {
			Active *bool   `json:"active,omitempty"`
			Id     *string `json:"id,omitempty"`
			Image  *string `json:"image,omitempty"`
			Name   *string `json:"name,omitempty"`
		} `json:"devices,omitempty"`
		HiSpeedMode *bool `json:"hi_speed_mode,omitempty"`
	} `json:"mornin_plus,omitempty"`
	Nickname *string `json:"nickname,omitempty"`
	QrioLock *struct {
		BdAddress *string `json:"bd_address,omitempty"`
		Device    *struct {
			Id    *int32  `json:"id,omitempty"`
			Image *string `json:"image,omitempty"`
			Name  *string `json:"name,omitempty"`
		} `json:"device,omitempty"`
		IsAvailable  *bool   `json:"is_available,omitempty"`
		SubBdAddress *string `json:"sub_bd_address,omitempty"`
	} `json:"qrio_lock,omitempty"`
	Settings *struct {
		// Button Button. Specify 'power-off' always if you want the air conditioner powered off. Empty means powered on.
		Button *string `json:"button,omitempty"`

		// Dir AC air direction. Empty means automatic.
		Dir *string `json:"dir,omitempty"`

		// Dirh AC horizontal air direction.
		Dirh *string `json:"dirh,omitempty"`

		// Mode AC operation mode. The range of operation modes which the air conditioner accepts depends on the air conditioner model. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model.
		Mode *string `json:"mode,omitempty"`

		// Temp Temperature. The temperature in string format. The unit is described in Aircon object. The range of Temperatures which the air conditioner accepts depends on the air conditioner model and operation mode. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model and operation mode.
		Temp *string `json:"temp,omitempty"`

		// TempUnit Temperature unit. 'c' or 'f' or '' for unknown.
		TempUnit  *string    `json:"temp_unit,omitempty"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`

		// Vol AC air volume. Empty means automatic. Numbers express the amount of volume. The range of AirVolumes which the air conditioner accepts depends on the air conditioner model and operation mode. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model and operation mode.
		Vol *string `json:"vol,omitempty"`
	} `json:"settings,omitempty"`
	Signals *[]struct {
		Id    *string `json:"id,omitempty"`
		Image *string `json:"image,omitempty"`
		Name  *string `json:"name,omitempty"`
	} `json:"signals,omitempty"`
	SmartMeter *struct {
		EchonetliteProperties *[]struct {
			Epc       *int       `json:"epc,omitempty"`
			Name      *string    `json:"name,omitempty"`
			UpdatedAt *time.Time `json:"updated_at,omitempty"`
			Val       *string    `json:"val,omitempty"`
		} `json:"echonetlite_properties,omitempty"`
	} `json:"smart_meter,omitempty"`
	Tv *struct {
		Buttons *[]struct {
			Image *string `json:"image,omitempty"`
			Label *string `json:"label,omitempty"`
			Name  *string `json:"name,omitempty"`
		} `json:"buttons,omitempty"`
		Layout *[]struct {
			Buttons *[]string `json:"buttons,omitempty"`
			Type    *string   `json:"type,omitempty"`
		} `json:"layout,omitempty"`
		State *struct {
			Input *string `json:"input,omitempty"`
		} `json:"state,omitempty"`
	} `json:"tv,omitempty"`

	// Type Appliance types. AC, TV, LIGHT, etc.
	Type *string `json:"type,omitempty"`
}

// ApplianceResponses defines model for ApplianceResponses.
type ApplianceResponses = []ApplianceResponse

// CreateApplianceRequest defines model for CreateApplianceRequest.
type CreateApplianceRequest struct {
	Device *string `json:"device,omitempty"`

	// Image Basename of the image file included in the app. Ex: 'ico_ac_1'.
	Image *string `json:"image,omitempty"`

	// ModelType Enum of 'AC', 'TV', 'Light'
	ModelType *string `json:"model_type,omitempty"`

	// Nickname Appliance name.
	Nickname *string `json:"nickname,omitempty"`
}

// CreateSignalParameters defines model for CreateSignalParameters.
type CreateSignalParameters struct {
	// Image Basename of the image file included in the app. Ex: 'ico_io'.
	Image *string `json:"Image,omitempty"`

	// Message JSON serialized object describing infrared signals. Includes 'data', 'freq' and 'format' keys.
	Message *string `json:"Message,omitempty"`

	// Name Signal name.
	Name *string `json:"Name,omitempty"`
}

// DetectApplianceRequest defines model for DetectApplianceRequest.
type DetectApplianceRequest struct {
	Device *string `json:"Device,omitempty"`

	// Message JSON serialized object describing infrared signals. Includes 'data', 'freq' and 'format' keys.
	Message *string `json:"Message,omitempty"`
}

// DeviceResponse defines model for DeviceResponse.
type DeviceResponse struct {
	BtMacAddress    string     `json:"bt_mac_address" validate:"required"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	FirmwareVersion string     `json:"firmware_version" validate:"required"`
	HumidityOffset  float32    `json:"humidity_offset" validate:"min=-20,max=20"`
	Id              string     `json:"id" validate:"required"`
	MacAddress      string     `json:"mac_address" validate:"required"`
	Name            string     `json:"name" validate:"required"`

	// NewestEvents The SensorValue key means 'te' = temperature, 'hu' = humidity, 'il' = illumination, 'mo' = movement. The val of 'mo' is always 1 and when movement event is captured created_at is updated.
	NewestEvents *map[string]struct {
		CreatedAt *time.Time `json:"created_at,omitempty"`
		Val       *float32   `json:"val,omitempty"`
	} `json:"newest_events,omitempty"`
	Online            *bool      `json:"online,omitempty"`
	SerialNumber      string     `json:"serial_number" validate:"required"`
	TemperatureOffset float32    `json:"temperature_offset" validate:"min=-5,max=5"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
}

// DeviceResponses defines model for DeviceResponses.
type DeviceResponses = []DeviceResponse

// EchonetLiteApplianceResponse defines model for EchonetLiteApplianceResponse.
type EchonetLiteApplianceResponse struct {
	Appliances *[]struct {
		Device *struct {
			BtMacAddress      *string    `json:"bt_mac_address,omitempty"`
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			FirmwareVersion   *string    `json:"firmware_version,omitempty"`
			HumidityOffset    *float32   `json:"humidity_offset,omitempty"`
			Id                *string    `json:"id,omitempty"`
			MacAddress        *string    `json:"mac_address,omitempty"`
			Name              *string    `json:"name,omitempty"`
			SerialNumber      *string    `json:"serial_number,omitempty"`
			TemperatureOffset *float32   `json:"temperature_offset,omitempty"`
			UpdatedAt         *time.Time `json:"updated_at,omitempty"`
		} `json:"Device,omitempty"`
		Id         *string `json:"id,omitempty"`
		Nickname   *string `json:"nickname,omitempty"`
		Properties *[]struct {
			Epc       *string    `json:"epc,omitempty"`
			UpdatedAt *time.Time `json:"updated_at,omitempty"`
			Val       *string    `json:"val,omitempty"`
		} `json:"properties,omitempty"`
		Type *string `json:"type,omitempty"`
	} `json:"appliances,omitempty"`
}

// EmptyObject defines model for EmptyObject.
type EmptyObject struct {
	DummyField *string `json:"dummy_field,omitempty"`
}

// HomeInvite defines model for HomeInvite.
type HomeInvite struct {
	Home *struct {
		Id   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	} `json:"home,omitempty"`
	Url  *string `json:"url,omitempty"`
	User *struct {
		Id       *string `json:"id,omitempty"`
		Nickname *string `json:"nickname,omitempty"`
	} `json:"user,omitempty"`
}

// HomeParams defines model for HomeParams.
type HomeParams struct {
	Name *string `json:"name,omitempty"`
}

// HomeResponse defines model for HomeResponse.
type HomeResponse struct {
	Id       *string `json:"id,omitempty"`
	Location *struct {
		Latitude  *float64 `json:"latitude,omitempty"`
		Longitude *float64 `json:"longitude,omitempty"`
		Radius    *float32 `json:"radius,omitempty"`
	} `json:"location,omitempty"`
	Name *string `json:"name,omitempty"`
	Town *struct {
		Id         *int32   `json:"id,omitempty"`
		Latitude   *float64 `json:"latitude,omitempty"`
		Longitude  *float64 `json:"longitude,omitempty"`
		Name       *string  `json:"name,omitempty"`
		Origin     *string  `json:"origin,omitempty"`
		Prefecture *string  `json:"prefecture,omitempty"`
	} `json:"town,omitempty"`
	Users *[]struct {
		Id            *string    `json:"id,omitempty"`
		JoinedAt      *time.Time `json:"joined_at,omitempty"`
		LocationState *string    `json:"location_state,omitempty"`
		Nickname      *string    `json:"nickname,omitempty"`
		Role          *string    `json:"role,omitempty"`
	} `json:"users,omitempty"`
}

// HomeResponses defines model for HomeResponses.
type HomeResponses = []HomeResponse

// HomeUserParams defines model for HomeUserParams.
type HomeUserParams struct {
	User *string `json:"user,omitempty"`
}

// HumidityOffsetParams defines model for HumidityOffsetParams.
type HumidityOffsetParams struct {
	// Offset Humidity offset value added to the measured humidity.
	Offset *float32 `json:"offset,omitempty" validate:"min=-20,max=20"`
}

// LightParams defines model for LightParams.
type LightParams struct {
	// Button Button name.
	Button *string `json:"button,omitempty"`
}

// LightProjectorParams defines model for LightProjectorParams.
type LightProjectorParams struct {
	// Button Button name.
	Button *string `json:"button,omitempty"`
}

// LightState defines model for LightState.
type LightState struct {
	Brightness *string `json:"brightness,omitempty"`
	LastButton *string `json:"last_button,omitempty"`
	Power      *string `json:"power,omitempty"`
}

// RefreshELPropertyRequest defines model for RefreshELPropertyRequest.
type RefreshELPropertyRequest struct {
	// EPC Comma separated EPCs in hex. eg: cf,da
	EPC *string `json:"EPC,omitempty"`
}

// ReorderAppliancesParams defines model for ReorderAppliancesParams.
type ReorderAppliancesParams struct {
	// Appliances List of all appliance IDs, comma separated.
	Appliances *string `json:"Appliances,omitempty"`
}

// ReorderSignalsParams defines model for ReorderSignalsParams.
type ReorderSignalsParams struct {
	// Signals List of all signal IDs, comma separated.
	Signals *string `json:"Signals,omitempty"`
}

// SendSignalDeviceParameters defines model for SendSignalDeviceParameters.
type SendSignalDeviceParameters struct {
	Message *string `json:"Message,omitempty"`
}

// SetELPropertyRequest defines model for SetELPropertyRequest.
type SetELPropertyRequest struct {
	// EPC EPC in hex string. eg: cf
	EPC *string `json:"EPC,omitempty"`

	// Val Value in hex string. String length must be 2x the number of bytes according to ECHONET Lite spec, and filled with zero if necessary. eg: 000000FF
	Val *string `json:"Val,omitempty"`
}

// Signal defines model for Signal.
type Signal struct {
	Id    *string `json:"id,omitempty"`
	Image *string `json:"image,omitempty"`
	Name  *string `json:"name,omitempty"`
}

// Signals defines model for Signals.
type Signals = []Signal

// TVParams defines model for TVParams.
type TVParams struct {
	// Button Button name.
	Button *string `json:"button,omitempty"`
}

// TVState defines model for TVState.
type TVState struct {
	Input *string `json:"input,omitempty"`
}

// TemperatureOffsetParams defines model for TemperatureOffsetParams.
type TemperatureOffsetParams struct {
	// Offset Temperature offset value added to the measured temperature.
	Offset *float32 `json:"offset,omitempty" validate:"min=-5,max=5"`
}

// Template defines model for Template.
type Template struct {
	Image     *string     `json:"image,omitempty"`
	Label     *string     `json:"label,omitempty"`
	Name      *string     `json:"name,omitempty"`
	Templates *[]Template `json:"templates,omitempty"`
	Text      *string     `json:"text,omitempty"`
	Type      *string     `json:"type,omitempty"`
	Uuid      *string     `json:"uuid,omitempty"`
	XSize     *int        `json:"x_size,omitempty"`
	YSize     *int        `json:"y_size,omitempty"`
}

// TransferRequest defines model for TransferRequest.
type TransferRequest struct {
	Devices *string `json:"Devices,omitempty"`
}

// UpdateDeviceParameters defines model for UpdateDeviceParameters.
type UpdateDeviceParameters struct {
	Name string `json:"name" validate:"required"`
}

// UpdateProfileParam defines model for UpdateProfileParam.
type UpdateProfileParam struct {
	Country      *UpdateProfileParamCountry      `json:"country,omitempty"`
	DistanceUnit *UpdateProfileParamDistanceUnit `json:"distance_unit,omitempty"`
	Nickname     *string                         `json:"nickname,omitempty"`
	TempUnit     *UpdateProfileParamTempUnit     `json:"temp_unit,omitempty"`
}

// UpdateProfileParamCountry defines model for UpdateProfileParam.Country.
type UpdateProfileParamCountry string

// UpdateProfileParamDistanceUnit defines model for UpdateProfileParam.DistanceUnit.
type UpdateProfileParamDistanceUnit string

// UpdateProfileParamTempUnit defines model for UpdateProfileParam.TempUnit.
type UpdateProfileParamTempUnit string

// UpdateSignalParameters defines model for UpdateSignalParameters.
type UpdateSignalParameters struct {
	// Image Basename of the image file included in the app. Ex: 'ico_io'.
	Image *string `json:"Image,omitempty"`

	// Name Signal name.
	Name *string `json:"Name,omitempty"`
}

// UserAndRole defines model for UserAndRole.
type UserAndRole struct {
	Role *string `json:"role,omitempty"`
	User *struct {
		Id       *string `json:"id,omitempty"`
		Nickname *string `json:"nickname,omitempty"`
	} `json:"user,omitempty"`
}

// UserAndRoles defines model for UserAndRoles.
type UserAndRoles = []UserAndRole

// UserResponse defines model for UserResponse.
type UserResponse struct {
	Id       string `json:"id" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
}

// Post1ApplianceOrdersFormdataRequestBody defines body for Post1ApplianceOrders for application/x-www-form-urlencoded ContentType.
type Post1ApplianceOrdersFormdataRequestBody = ReorderAppliancesParams

// Post1AppliancesFormdataRequestBody defines body for Post1Appliances for application/x-www-form-urlencoded ContentType.
type Post1AppliancesFormdataRequestBody = CreateApplianceRequest

// Post1AppliancesApplianceidFormdataRequestBody defines body for Post1AppliancesApplianceid for application/x-www-form-urlencoded ContentType.
type Post1AppliancesApplianceidFormdataRequestBody = ApplianceParams

// Post1AppliancesApplianceidAirconSettingsFormdataRequestBody defines body for Post1AppliancesApplianceidAirconSettings for application/x-www-form-urlencoded ContentType.
type Post1AppliancesApplianceidAirconSettingsFormdataRequestBody = AirConParams

// Post1AppliancesApplianceidDeleteFormdataRequestBody defines body for Post1AppliancesApplianceidDelete for application/x-www-form-urlencoded ContentType.
type Post1AppliancesApplianceidDeleteFormdataRequestBody = EmptyObject

// Post1AppliancesApplianceidLightFormdataRequestBody defines body for Post1AppliancesApplianceidLight for application/x-www-form-urlencoded ContentType.
type Post1AppliancesApplianceidLightFormdataRequestBody = LightParams

// Post1AppliancesApplianceidLightProjectorFormdataRequestBody defines body for Post1AppliancesApplianceidLightProjector for application/x-www-form-urlencoded ContentType.
type Post1AppliancesApplianceidLightProjectorFormdataRequestBody = LightProjectorParams

// Post1AppliancesApplianceidSignalOrdersFormdataRequestBody defines body for Post1AppliancesApplianceidSignalOrders for application/x-www-form-urlencoded ContentType.
type Post1AppliancesApplianceidSignalOrdersFormdataRequestBody = ReorderSignalsParams

// Post1AppliancesApplianceidSignalsFormdataRequestBody defines body for Post1AppliancesApplianceidSignals for application/x-www-form-urlencoded ContentType.
type Post1AppliancesApplianceidSignalsFormdataRequestBody = CreateSignalParameters

// Post1AppliancesApplianceidTvFormdataRequestBody defines body for Post1AppliancesApplianceidTv for application/x-www-form-urlencoded ContentType.
type Post1AppliancesApplianceidTvFormdataRequestBody = TVParams

// Post1DetectapplianceFormdataRequestBody defines body for Post1Detectappliance for application/x-www-form-urlencoded ContentType.
type Post1DetectapplianceFormdataRequestBody = DetectApplianceRequest

// Post1DevicesDeviceidFormdataRequestBody defines body for Post1DevicesDeviceid for application/x-www-form-urlencoded ContentType.
type Post1DevicesDeviceidFormdataRequestBody = UpdateDeviceParameters

// Post1DevicesDeviceidDeleteFormdataRequestBody defines body for Post1DevicesDeviceidDelete for application/x-www-form-urlencoded ContentType.
type Post1DevicesDeviceidDeleteFormdataRequestBody = EmptyObject

// Post1DevicesDeviceidHumidityOffsetFormdataRequestBody defines body for Post1DevicesDeviceidHumidityOffset for application/x-www-form-urlencoded ContentType.
type Post1DevicesDeviceidHumidityOffsetFormdataRequestBody = HumidityOffsetParams

// Post1DevicesDeviceidSendFormdataRequestBody defines body for Post1DevicesDeviceidSend for application/x-www-form-urlencoded ContentType.
type Post1DevicesDeviceidSendFormdataRequestBody = SendSignalDeviceParameters

// Post1DevicesDeviceidTemperatureOffsetFormdataRequestBody defines body for Post1DevicesDeviceidTemperatureOffset for application/x-www-form-urlencoded ContentType.
type Post1DevicesDeviceidTemperatureOffsetFormdataRequestBody = TemperatureOffsetParams

// Post1EchonetliteAppliancesApplianceidRefreshFormdataRequestBody defines body for Post1EchonetliteAppliancesApplianceidRefresh for application/x-www-form-urlencoded ContentType.
type Post1EchonetliteAppliancesApplianceidRefreshFormdataRequestBody = RefreshELPropertyRequest

// Post1EchonetliteAppliancesApplianceidSetFormdataRequestBody defines body for Post1EchonetliteAppliancesApplianceidSet for application/x-www-form-urlencoded ContentType.
type Post1EchonetliteAppliancesApplianceidSetFormdataRequestBody = SetELPropertyRequest

// Post1HomesFormdataRequestBody defines body for Post1Homes for application/x-www-form-urlencoded ContentType.
type Post1HomesFormdataRequestBody = HomeParams

// Post1HomesHomeidFormdataRequestBody defines body for Post1HomesHomeid for application/x-www-form-urlencoded ContentType.
type Post1HomesHomeidFormdataRequestBody = HomeParams

// Post1HomesHomeidKickFormdataRequestBody defines body for Post1HomesHomeidKick for application/x-www-form-urlencoded ContentType.
type Post1HomesHomeidKickFormdataRequestBody = HomeUserParams

// Post1HomesHomeidOwnerFormdataRequestBody defines body for Post1HomesHomeidOwner for application/x-www-form-urlencoded ContentType.
type Post1HomesHomeidOwnerFormdataRequestBody = HomeUserParams

// Post1HomesHomeidTransferTohomeidFormdataRequestBody defines body for Post1HomesHomeidTransferTohomeid for application/x-www-form-urlencoded ContentType.
type Post1HomesHomeidTransferTohomeidFormdataRequestBody = TransferRequest

// Post1SignalsSignalidFormdataRequestBody defines body for Post1SignalsSignalid for application/x-www-form-urlencoded ContentType.
type Post1SignalsSignalidFormdataRequestBody = UpdateSignalParameters

// Post1SignalsSignalidDeleteFormdataRequestBody defines body for Post1SignalsSignalidDelete for application/x-www-form-urlencoded ContentType.
type Post1SignalsSignalidDeleteFormdataRequestBody = EmptyObject

// Post1SignalsSignalidSendFormdataRequestBody defines body for Post1SignalsSignalidSend for application/x-www-form-urlencoded ContentType.
type Post1SignalsSignalidSendFormdataRequestBody = EmptyObject

// Post1UsersMeFormdataRequestBody defines body for Post1UsersMe for application/x-www-form-urlencoded ContentType.
type Post1UsersMeFormdataRequestBody = UpdateProfileParam
