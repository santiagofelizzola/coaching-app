package models

type User struct {
	BaseModel
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"` // In production, this should be hashed!
	Role     string `json:"role"`     // "coach" or "player"

	Teams []Team `gorm:"many2many:user_teams;"`
	Tasks []Task `gorm:"many2many:user_tasks;"`
}
