package factories

import (
	"github.com/jaswdr/faker"
	"gorm.io/gorm"

	"github.com/yugarinn/catapi.cat/app/breeds/models"
	"github.com/yugarinn/catapi.cat/app/breeds/inputs"
	"github.com/yugarinn/catapi.cat/connections"
)


var database *gorm.DB = connections.Database()

func CreateBreed(properties inputs.CreateBreedInput) breeds.Breed {
    breed := breeds.Breed{Uuid: properties.Uuid, Name: properties.Name, Description: properties.Description}

	fake := faker.New()

	if len(breed.Uuid) == 0 {
		breed.Uuid = fake.UUID().V4()
	}

	if len(breed.Name) == 0 {
		breed.Name = fake.Person().Name()
	}

	if len(breed.Description) == 0 {
		breed.Description = fake.Person().Contact().Email
	}

	database.Create(&breed)

	return breed
}

func CreateBreedsList(number int, properties inputs.CreateBreedInput) []breeds.Breed {
	var breeds []breeds.Breed

	for i := 0; i < number; i++ {
		breeds = append(breeds, CreateBreed(properties))
	}

	return breeds
}
