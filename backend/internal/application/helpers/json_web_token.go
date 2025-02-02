package helpers

import (
	"backend/config"
	"backend/internal/application/typings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(claims jwt.MapClaims) Either[string] {
	var either Either[string]

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(config.Env.JWT_SECRET)

	if err != nil {
		either.Left(typings.Error{
			StatusCode: 500,
			Message:    "Failed to generate JSON Web Token",
			Reason:     err.Error(),
		})
	} else {
		either.Right(tokenString)
	}

	return either
}
