package facts

import (
	"gorm.io/gorm"
)

type Fact struct {
	gorm.Model
	Uuid string
	Text string
}

func (Fact) TableName() string {
	return "facts_facts"
}
