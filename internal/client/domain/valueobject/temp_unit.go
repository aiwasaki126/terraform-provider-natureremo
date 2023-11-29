package valueobject

import "fmt"

// TempUnit is the temperature unit setting in profile.
// TempUnit must be one of c and f. Each of c, f represents Celsius, Fahrenheit.
type TempUnit string

// NewTempUnit is the constructor for TempUnit.
// Invalid temperature unit input will return error.
func NewTempUnit(t string) (TempUnit, error) {
	switch t {
	case "c":
		break
	case "f":
		break
	default:
		return "", fmt.Errorf("invalid temperature unit %s", t)
	}
	return TempUnit(t), nil
}

// Value returns temperature unit in string.
func (t TempUnit) Value() string {
	return string(t)
}
