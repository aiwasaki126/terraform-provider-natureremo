package valueobject

import "fmt"

// DistanceUnit is the distance unit setting in profile.
// DistanceUnit must be one of metric and imperial.
type DistanceUnit string

// NewDistanceUnit is the constructor for DistanceUnit.
// Invalid distance unit input will return error.
func NewDistanceUnit(d string) (DistanceUnit, error) {
	switch d {
	case "metric":
		break
	case "imperial":
		break
	default:
		return "", fmt.Errorf("invalid distance unit %s", d)
	}
	return DistanceUnit(d), nil
}

// Value returns distance unit in string.
func (d DistanceUnit) Value() string {
	return string(d)
}
