package usecase

import (
	"context"
	"terraform-provider-natureremo/internal/client/domain/repository"
	"terraform-provider-natureremo/internal/client/domain/valueobject"

	"github.com/go-playground/validator/v10"
)

// UpdateDevice is the use case of updating the device.
type UpdateDevice struct {
	repositoryDevice repository.Device
}

func NewUpdateDevice(r repository.Device) *UpdateDevice {
	return &UpdateDevice{
		repositoryDevice: r,
	}
}

func (u *UpdateDevice) UpdateDevice(ctx context.Context, deviceDto DeviceDto) (*DeviceDto, error) {
	if err := validator.New().Struct(deviceDto); err != nil {
		return nil, err
	}
	d, err := u.repositoryDevice.GetDevice(ctx, deviceDto.Id)
	if err != nil {
		return nil, err
	}
	newSetting, err := valueobject.NewDeviceSetting(deviceDto.Name, deviceDto.HumidityOffset, deviceDto.TemperatureOffset)
	if err != nil {
		return nil, err
	}
	if err := d.UpdateSetting(*newSetting); err != nil {
		return nil, err
	}
	d, err = u.repositoryDevice.UpdateDevice(ctx, d)
	if err != nil {
		return nil, err
	}
	return newDeviceDtoFromEntity(d), nil
}
