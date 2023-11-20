package valueobject

import "fmt"

type Country string

func NewCountry(c string) (Country, error) {
	switch c {
	case "JP":
		break
	case "US":
		break
	case "CA":
		break
	case "SG":
		break
	case "AU":
		break
	case "NZ":
		break
	case "OTHERS":
		break
	default:
		return "", fmt.Errorf("invalid country %s", c)
	}
	return Country(c), nil
}

func (c Country) Value() string {
	return string(c)
}
