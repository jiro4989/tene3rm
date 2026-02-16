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
