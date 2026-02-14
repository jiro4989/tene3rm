package domain

import "fmt"

type RangeInt struct {
	value, minValue, maxValue int
}

func NewRangeInt(value, mn, mx int) (RangeInt, error) {
	if mx < mn {
		return RangeInt{}, fmt.Errorf("'min' must be less than 'max': min = %d, max = %d", mn, mx)
	}

	if !(mn <= value && value <= mx) {
		return RangeInt{}, fmt.Errorf("'value' between 'min' and 'max': value = %d, min = %d, max = %d", value, mn, mx)
	}

	return RangeInt{
		value:    value,
		minValue: mn,
		maxValue: mx,
	}, nil
}

func (r RangeInt) SafePlus(n int) RangeInt {
	v := r.value + n

	if r.minValue <= v && v <= r.maxValue {
		r.value = v
	} else if v < r.minValue {
		r.value = r.minValue
	} else if r.maxValue < v {
		r.value = r.maxValue
	}

	return r
}

func (r RangeInt) Value() int {
	return r.value
}
