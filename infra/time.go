package infra

import "time"

type TimeGenerator interface {
	Now() time.Time
}

type TimeGeneratorImpl struct{}

func NewTimeGeneratorImpl() TimeGeneratorImpl {
	return TimeGeneratorImpl{}
}

func (t TimeGeneratorImpl) Now() time.Time {
	return time.Now()
}

type MockTimeGenerator struct {
	t time.Time
}

func NewMockTimeGenerator(t time.Time) MockTimeGenerator {
	return MockTimeGenerator{
		t: t,
	}
}

func (t MockTimeGenerator) Now() time.Time {
	return t.t
}
