package usecase

import (
	"context"
	"terraform-provider-natureremo/internal/client/domain/repository"
)

type GetDevice struct {
	repositoryDevice repository.Device
}

func NewGetDevice(r repository.Device) *GetDevice {
	return &GetDevice{
		repositoryDevice: r,
	}
}

func (u *GetDevice) GetDevice(ctx context.Context, id string) (*DeviceDto, error) {
	device, err := u.repositoryDevice.GetDevice(ctx, id)
	if err != nil {
		return nil, err
	}
	return newDeviceDtoFromEntity(device), nil
}
