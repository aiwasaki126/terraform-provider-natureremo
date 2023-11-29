package valueobject

import "fmt"

// Country is the country setting in profile.
// Country must be one of JP, US, CA, SG, AU, NZ and OTHERS.
type Country string

// NewCountry is the constructor for Country.
// Invalid country name input will return error.
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

// Value returns country in string.
func (c Country) Value() string {
	return string(c)
}
