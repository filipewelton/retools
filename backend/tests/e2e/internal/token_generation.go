package internal

import (
	"backend/internal/application/helpers"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAPICredential() string {
	token := helpers.GenerateJWT(jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	if token.IsLeft {
		panic(token.LeftResponse)
	}

	return token.RightResponse
}
