package responses

import (
	models "github.com/yugarinn/pigeon-api/app/users/models"
)

type UserResponse struct {
	ID uint64          `json:"id"`
	Email string       `json:"email"`
	Name string        `json:"name"`
	CountryCode string `json:"countryCode"`
	PhoneNumber string `json:"phoneNumber"`
}

type PaginatedUserResponse struct {
	Users []UserResponse  `json:"users"`
	Pagination Pagination `json:"pagination"`
}

func SerializeUser(user models.User) UserResponse {
	return UserResponse{
		ID: user.ID,
		Email: user.Email,
		Name: user.Name,
		CountryCode: user.CountryCode,
		PhoneNumber: user.PhoneNumber,
	}
}
