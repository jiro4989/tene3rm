package infra

import "time"

type StateDTO struct {
	FailCount int
	Created   *time.Time
}
