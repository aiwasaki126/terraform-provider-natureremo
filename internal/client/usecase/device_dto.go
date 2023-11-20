package usecase

import (
	"terraform-provider-natureremo/internal/client/domain/entity"

	"github.com/go-playground/validator/v10"
)

type DeviceDto struct {
	Id string `json:"id" validate:"required"`
	// setting items
	Name              string  `json:"name" validate:"required"`
	HumidityOffset    int64   `json:"humidity_offset" validate:"min=-20,max=20"`
	TemperatureOffset float64 `json:"temperature_offset" validate:"min=-5,max=5"`
	// spec items
	SerialNumber    string `json:"serial_number"`
	MacAddress      string `json:"mac_address"`
	BtMacAddress    string `json:"bt_mac_address"`
	FirmwareVersion string `json:"firmware_version"`

	// Users Deprecated. Do not use in new code.
	Users []struct {
		Id       string `json:"id"`
		Nickname string `json:"nickname"`
	} `json:"users"`
}

func newDeviceDtoFromEntity(d *entity.Device) *DeviceDto {
	users := d.GetUsers()

	usersDto := make([]struct {
		Id       string
		Nickname string
	}, 0, len(users))
	for _, u := range users {
		usersDto = append(usersDto, struct {
			Id       string
			Nickname string
		}{u.GetId(), u.GetNickname()})
	}
	dto := &DeviceDto{
		Id:                d.GetId(),
		Name:              d.GetName(),
		HumidityOffset:    d.GetHumidityOffset(),
		TemperatureOffset: d.GetTemperatureOffset(),
		SerialNumber:      d.GetSerialNumber(),
		MacAddress:        d.GetMacAddress(),
		BtMacAddress:      d.GetBtMacAddress(),
		FirmwareVersion:   d.GetFirmwareVersion(),
		Users: []struct {
			Id       string `json:"id"`
			Nickname string `json:"nickname"`
		}(usersDto),
	}
	return dto
}

func NewDeviceDto(id string, name string, humidityOffset int64, temperatureOffset float64) (*DeviceDto, error) {
	dto := &DeviceDto{
		Id:                id,
		Name:              name,
		HumidityOffset:    humidityOffset,
		TemperatureOffset: temperatureOffset,
	}
	if err := validator.New().Struct(dto); err != nil {
		return nil, err
	}
	return dto, nil
}
