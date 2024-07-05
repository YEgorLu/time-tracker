package models

import "github.com/google/uuid"

type TimerStartRes struct {
	TimerId uuid.UUID `json:"timerId"`
}
