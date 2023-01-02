package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	managers "github.com/yugarinn/catapi.cat/app/breeds/managers"
	responses "github.com/yugarinn/catapi.cat/http/responses"
)


func GetBreeds(context *gin.Context) {
	breeds := managers.GetPaginatedBreeds(1, DEFAULT_PAGE_SIZE)

	response := []responses.BreedResponse{}

	for _, breed := range breeds {
		response = append(response, responses.SerializeBreed(breed))
	}

	context.JSON(http.StatusOK, response)
}

func GetBreed(context *gin.Context) {
	breedId := context.Param("breedId")
	breed, retrievalError := managers.GetBreed(breedId)

	if retrievalError != nil {
		HttpErrorHandler(context, 404, "breed not found")
		return
	}

	context.JSON(http.StatusOK, responses.SerializeBreed(breed))
}
