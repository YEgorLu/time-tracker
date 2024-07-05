package models

import "github.com/google/uuid"

type TaskWorktime struct {
	TaskId uuid.UUID
	Ms     int64
}
