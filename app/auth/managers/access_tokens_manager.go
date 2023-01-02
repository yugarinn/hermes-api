package auth

import (
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/yugarinn/catapi.cat/connections"

	auth "github.com/yugarinn/catapi.cat/app/auth/models"
	users "github.com/yugarinn/catapi.cat/app/users/models"

	"github.com/golang-jwt/jwt/v4"
)


var database *gorm.DB = connections.Database()

type JWTClaims struct {
    UserID    string `json:"userId"`
    UserEmail string `json:"userEmail"`
    jwt.StandardClaims
}

func GenerateAccessTokenForUser(user users.User) (auth.AccessToken, error) {
	expirationDate := generateExpirationDate()

	jwtClaims := JWTClaims{
        UserEmail: user.Email,
        UserID: strconv.FormatUint(user.ID, 10),
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
        },
	}

	secret := os.Getenv("JWT_SECRET_KEY")

	token, _ := generateJWT(jwtClaims, secret)
    accessToken := auth.AccessToken{UserId: user.ID, Token: token, ExpiresAt: expirationDate}
	result := database.Create(&accessToken)

	return accessToken, result.Error
}

func GetAccessTokenByBy(field string, value any) (auth.AccessToken, error) {
	var accessToken auth.AccessToken
	result := database.Where(field, value).First(&accessToken)

	return accessToken, result.Error
}

func generateExpirationDate() time.Time {
	return time.Now().Add(time.Hour * 24)
}


func generateJWT(claims JWTClaims, secret string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString([]byte(secret))
}
