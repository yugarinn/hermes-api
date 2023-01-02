package responses

import (
	models "github.com/yugarinn/catapi.cat/app/breeds/models"
)

type BreedResponse struct {
	ID string          `json:"id"`
	Name string        `json:"name"`
	Description string `json:"description"`
}

func SerializeBreed(breed models.Breed) BreedResponse {
	return BreedResponse{
		ID: breed.Uuid,
		Name: breed.Name,
		Description: breed.Description,
	}
}
