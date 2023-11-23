package entity

import "terraform-provider-natureremo/internal/client/domain/valueobject"

type Device struct {
	id string
	valueobject.DeviceSetting
	valueobject.DeviceSpec
}

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
