package breeds

import (
	"gorm.io/gorm"

	"github.com/yugarinn/catapi.cat/connections"
	"github.com/yugarinn/catapi.cat/app/breeds/models"
)


var database *gorm.DB = connections.Database()

func GetPaginatedBreeds(page int, size int) []breeds.Breed {
	var breeds []breeds.Breed
	database.Find(&breeds)

	return breeds
}

func GetBreed(uuid string) (breeds.Breed, error) {
	var breed breeds.Breed
	result := database.Where("uuid", uuid).First(&breed)

	return breed, result.Error
}
