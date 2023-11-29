package valueobject

// DeviceSpec is the device spec of Nature Remo.
type DeviceSpec struct {
	serialNumber    string
	btMacAddress    string
	firmWareVersion string
	macAddress      string
}

// NewDeviceSpec is the constructor for DeviceSpec.
func NewDeviceSpec(serialNumber, btMacAddress, firmWareVersion, macAddress string) *DeviceSpec {
	return &DeviceSpec{
		serialNumber:    serialNumber,
		btMacAddress:    btMacAddress,
		firmWareVersion: firmWareVersion,
		macAddress:      macAddress,
	}
}

func (s *DeviceSpec) GetSerialNumber() string {
	return s.serialNumber
}

func (s *DeviceSpec) GetBtMacAddress() string {
	return s.btMacAddress
}

func (s *DeviceSpec) GetFirmwareVersion() string {
	return s.firmWareVersion
}

func (s *DeviceSpec) GetMacAddress() string {
	return s.macAddress
}
