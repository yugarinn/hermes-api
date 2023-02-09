package users

import (
	"time"

	"gorm.io/gorm"
)


type UserValidationCode struct {
	gorm.Model
	ID uint64
	UserId uint64
	Code string
	IsUsed bool
	ExpiresAt time.Time
}

func (UserValidationCode) TableName() string {
	return "users_validation_codes"
}
