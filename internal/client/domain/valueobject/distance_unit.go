package valueobject

import "fmt"

type DistanceUnit string

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

func (d DistanceUnit) Value() string {
	return string(d)
}
