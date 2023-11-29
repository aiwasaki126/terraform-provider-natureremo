package usecase

import (
	"context"
	"terraform-provider-natureremo/internal/client/domain/repository"
)

// DeleteDevice is the use case of deleting the device.
type DeleteDevice struct {
	repositoryDevice repository.Device
}

func NewDeleteDevice(r repository.Device) *DeleteDevice {
	return &DeleteDevice{
		repositoryDevice: r,
	}
}

func (u *DeleteDevice) DeleteDevice(ctx context.Context, id string) error {
	d, err := u.repositoryDevice.GetDevice(ctx, id)
	if err != nil {
		return err
	}
	if err := u.repositoryDevice.DeleteDevice(ctx, d); err != nil {
		return err
	}
	return nil
}
