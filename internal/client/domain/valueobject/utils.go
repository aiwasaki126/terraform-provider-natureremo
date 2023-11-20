package valueobject

import (
	"fmt"
	"math"
)

type valueRange struct {
	minValue float64
	maxValue float64
}

func newValueRange(min, max float64) (*valueRange, error) {
	if min >= max {
		return nil, fmt.Errorf("min must be less than max")
	}
	return &valueRange{minValue: min, maxValue: max}, nil
}

func (r *valueRange) min() float64 {
	return r.minValue
}

func (r *valueRange) max() float64 {
	return r.maxValue
}

func hasValidIncrementedValue(v, increment float64, valueRange *valueRange) error {
	if increment <= 0 {
		return fmt.Errorf("invalid increment")
	}
	if v < valueRange.min() || valueRange.max() < v {
		return fmt.Errorf("invalid offset in range")
	}
	if math.Mod(v, increment) != 0 {
		return fmt.Errorf("invalid offset in range increment")
	}
	return nil
}
