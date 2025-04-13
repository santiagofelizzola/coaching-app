package models

import (
	"time"
)

type Task struct {
	BaseModel
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	CreatedByID uint
	CreatedBy   User

	Users []User `gorm:"many2many:user_tasks;"`
}
