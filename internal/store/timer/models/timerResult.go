package models

import (
	"time"

	"github.com/google/uuid"
)

type TimerResult struct {
	Id             uuid.UUID
	TaskId         uuid.UUID
	TimestampStart time.Time
}

func (r TimerResult) Elapsed() time.Duration {
	return time.Now().Sub(r.TimestampStart)
}
