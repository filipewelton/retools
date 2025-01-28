package internal

import (
	"backend/config"

	"github.com/go-chi/jwtauth"
)

func GenerateToken() string {
	tokenSettings := jwtauth.New("HS256", config.Env.JWT_SECRET, nil)

	_, tokenString, err := tokenSettings.Encode(map[string]interface{}{})

	if err != nil {
		panic(err)
	}

	return tokenString
}
