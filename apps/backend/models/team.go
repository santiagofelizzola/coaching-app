package models

type Team struct {
	BaseModel
	Name  string
	Users []User `gorm:"many2many:user_teams;"`
}
