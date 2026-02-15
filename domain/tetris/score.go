package tetris

import "math"

type Score struct {
	value    int
	maxValue int
}

func newScore(value, maxValue int) Score {
	return Score{
		value:    value,
		maxValue: maxValue,
	}
}

func newDefaultScore() Score {
	return newScore(0, 9999)
}

func (s Score) Plus() Score {
	s.value = int(math.Min(float64(s.value+100), float64(s.maxValue)))
	return s
}
