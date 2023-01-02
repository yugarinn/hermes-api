package breeds

import (
	"gorm.io/gorm"

	"github.com/yugarinn/catapi.cat/connections"
	"github.com/yugarinn/catapi.cat/app/facts/models"
)


var database *gorm.DB = connections.Database()

func GetPaginatedFacts(page int, size int) []facts.Fact {
	var facts []facts.Fact
	database.Find(&facts)

	return facts
}
