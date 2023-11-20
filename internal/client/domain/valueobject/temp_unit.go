package valueobject

import "fmt"

type TempUnit string

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

func (t TempUnit) Value() string {
	return string(t)
}
