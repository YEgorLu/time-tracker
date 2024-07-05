package models

import (
	"time"

	"github.com/google/uuid"
)

type WorkTime struct {
	Id        uuid.UUID
	TaskId    uuid.UUID
	Ms        int64
	Timestamp time.Time
}
