package models

import "github.com/google/uuid"

type TaskWorktime struct {
	TaskId uuid.UUID `json:"taskId"`
	Ms     int64     `json:"ms"`
}
