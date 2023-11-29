package entity

import "terraform-provider-natureremo/internal/client/domain/valueobject"

// Device is the device entity such as Nature Remo or Nature Remo Nano.
type Device struct {
	id string
	valueobject.DeviceSetting
	valueobject.DeviceSpec
}

// ReconstructDevice is the entity constructor from data in DB.
// This constructor is assumed to be used in repository.
func ReconstructDevice(id string, setting valueobject.DeviceSetting, spec valueobject.DeviceSpec) *Device {
	return &Device{
		id:            id,
		DeviceSetting: setting,
		DeviceSpec:    spec,
	}
}

func (d *Device) GetId() string {
	return d.id
}

func (d *Device) GetSetting() valueobject.DeviceSetting {
	return d.DeviceSetting
}

func (d *Device) UpdateSetting(s valueobject.DeviceSetting) error {
	d.DeviceSetting = s
	return nil
}

func (d *Device) GetSpec() valueobject.DeviceSpec {
	return d.DeviceSpec
}
