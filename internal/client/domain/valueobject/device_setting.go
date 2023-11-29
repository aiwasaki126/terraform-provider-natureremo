package valueobject

// DeviceSetting is the device setting of Nature Remo.
type DeviceSetting struct {
	name              string
	humidityOffset    int64
	temperatureOffset float64
}

// NewDeviceSetting is the constructor for DeviceSetting.
// Setting values have restrictions as follows.
// Humidity offset ranges -20 to 20 in increments of 5.
// Temperature offset ranges -5 to 5 in increments of 0.5.
func NewDeviceSetting(name string, humidityOffset int64, temperatureOffset float64) (*DeviceSetting, error) {
	// validation: humidity
	humidtyRange, err := newValueRange(-20, 20)
	if err != nil {
		return nil, err
	}
	var humidtyIncrement float64 = 5
	if err := hasValidIncrementedValue(float64(humidityOffset), humidtyIncrement, humidtyRange); err != nil {
		return nil, err
	}
	// validation: temperature
	tempRange, err := newValueRange(-5, 5)
	if err != nil {
		return nil, err
	}
	var tempIncrement float64 = 0.5
	if err := hasValidIncrementedValue(temperatureOffset, tempIncrement, tempRange); err != nil {
		return nil, err
	}
	return &DeviceSetting{
		name:              name,
		humidityOffset:    humidityOffset,
		temperatureOffset: temperatureOffset,
	}, nil
}

func (s *DeviceSetting) GetName() string {
	return s.name
}

func (s *DeviceSetting) GetHumidityOffset() int64 {
	return s.humidityOffset
}

func (s *DeviceSetting) GetTemperatureOffset() float64 {
	return s.temperatureOffset
}
