package responses

import (
	"time"

	models "github.com/yugarinn/hermes-api/app/auth/models"
)

type AccessTokenResponse struct {
	Token string `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func SerializeAccessToken(accessToken models.AccessToken) AccessTokenResponse {
	return AccessTokenResponse{
		Token: accessToken.Token,
		ExpiresAt: accessToken.ExpiresAt,
	}
}
