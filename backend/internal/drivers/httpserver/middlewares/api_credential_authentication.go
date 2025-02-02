package middlewares

import (
	"backend/config"
	"backend/internal/application/typings"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticateAPICredential(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			authorization = r.Header.Get("Authorization")
			apiCredential = strings.ReplaceAll(authorization, "Bearer ", "")
		)

		_, jwtErr := jwt.ParseWithClaims(
			apiCredential,
			&jwt.MapClaims{},
			func(t *jwt.Token) (interface{}, error) {
				if alg := t.Method.Alg(); alg != jwt.SigningMethodHS256.Name {
					return nil, errors.New("API credential is invalid")
				}

				exp, err := t.Claims.GetExpirationTime()

				if err != nil || exp == nil {
					return nil, errors.New("API credential is invalid")
				} else if now := time.Now().Unix(); now > exp.Unix() {
					return nil, errors.New("API credential is expired")
				}

				return []byte(config.Env.JWT_SECRET), nil
			})

		if jwtErr != nil {
			err := typings.Error{
				Message: "Unauthorized API access",
				Reason:  jwtErr.Error(),
			}

			render.Status(r, 401)
			render.JSON(w, r, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
