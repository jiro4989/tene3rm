package tetris

import "math"

type Score struct {
	value    int
	maxValue int
}

func newScore(value int) Score {
	return Score{
		value:    value,
		maxValue: 9999,
	}
}

func newDefaultScore() Score {
	return newScore(0)
}

func (s Score) Plus() Score {
	s.value = int(math.Min(float64(s.value+100), float64(s.maxValue)))
	return s
}

func (s Score) GreaterThan(s2 Score) bool {
	return s.value >= s2.value
}
