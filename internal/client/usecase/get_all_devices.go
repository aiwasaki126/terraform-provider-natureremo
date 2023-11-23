package usecase

import (
	"context"
	"terraform-provider-natureremo/internal/client/domain/repository"
)

type GetAllDevices struct {
	repositoryDevice repository.Device
}

func NewGetAllDevices(r repository.Device) *GetAllDevices {
	return &GetAllDevices{
		repositoryDevice: r,
	}
}

func (u *GetAllDevices) GetAllDevices(ctx context.Context) ([]*DeviceDto, error) {
	devices, err := u.repositoryDevice.GetAllDevices(ctx)
	if err != nil {
		return nil, err
	}
	deviceDtos := make([]*DeviceDto, 0, len(devices))
	for _, d := range devices {
		deviceDtos = append(deviceDtos, newDeviceDtoFromEntity(d))
	}
	return deviceDtos, nil
}
