package models

import "github.com/google/uuid"

type TimerStartReq struct {
	TaskId uuid.UUID `json:"taskId"`
	UserId uuid.UUID `json:"userId"`
}
