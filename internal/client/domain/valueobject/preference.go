package valueobject

// Preference is the user preference in profile.
// Preference is consist of country, distance unit and temperature unit.
type Preference struct {
	country      Country
	distanceUnit DistanceUnit
	tempUnit     TempUnit
}

// NewPreference is the constructor for the user preference.
func NewPreference(country, distanceUnit, tempUnit string) (*Preference, error) {
	c, err := NewCountry(country)
	if err != nil {
		return nil, err
	}
	d, err := NewDistanceUnit(distanceUnit)
	if err != nil {
		return nil, err
	}
	t, err := NewTempUnit(tempUnit)
	if err != nil {
		return nil, err
	}
	return &Preference{
		country:      c,
		distanceUnit: d,
		tempUnit:     t,
	}, nil
}

func (p *Preference) GetCountry() string {
	return p.country.Value()
}

func (p *Preference) GetDistanceUnit() string {
	return p.distanceUnit.Value()
}

func (p *Preference) GetTempUnit() string {
	return p.tempUnit.Value()
}
