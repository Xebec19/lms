package helpers

import (
	"time"

	"github.com/Xebec19/lms/common/utils"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(secret []byte, userId uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
		"iss": utils.GetConfig().APP_URL,
		"aud": utils.GetConfig().WEB_URL,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
