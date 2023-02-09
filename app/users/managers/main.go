package users

import (
	"gorm.io/gorm"

	"github.com/yugarinn/pigeon-api/connections"
)

var database *gorm.DB = connections.Database()
