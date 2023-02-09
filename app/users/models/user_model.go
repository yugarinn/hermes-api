package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint64
	Email string
	CountryCode string
	PhoneNumber string
	PhonePrefix string
	Name string
	IsConfirmed bool
}

func (User) TableName() string {
	return "users_users"
}
