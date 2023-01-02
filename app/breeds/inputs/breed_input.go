package inputs

type CreateBreedInput struct {
	Uuid string `json:"uuid"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
