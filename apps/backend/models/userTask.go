package models

import (
	"time"
)

type UserTask struct {
	BaseModel
	UserID      uint
	TaskID      uint
	Status      string
	CompletedAt *time.Time

	User User
	Task Task
}
