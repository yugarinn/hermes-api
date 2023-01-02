package breeds

import (
	"gorm.io/gorm"
)

type Breed struct {
	gorm.Model
	Uuid string
	Name string
	Description string
}

func (Breed) TableName() string {
	return "breeds_breeds"
}
