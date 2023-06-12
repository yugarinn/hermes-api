package users

import (
	"gorm.io/gorm"

	"github.com/yugarinn/hermes-api/connections"
)

var database *gorm.DB = connections.Database()
