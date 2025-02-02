package createsession

import (
	"backend/internal/application/contracts"
	"backend/internal/application/helpers"
	"backend/internal/application/typings"
	"backend/internal/domain/entities"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=48"`
}

type Params struct {
	UserRepository contracts.UserRepository
	Payload        Payload
}

type UserCredential struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Response struct {
	User       entities.MappedUserEntity `json:"user"`
	Credential UserCredential            `json:"credential"`
}

func Run(params Params) helpers.Either[Response] {
	var (
		either            helpers.Either[Response]
		payloadValidation = validatePayload(params.Payload)
	)

	if payloadValidation.IsLeft {
		either.Left(payloadValidation.LeftResponse)
		return either
	}

	var (
		email = params.Payload.Email
		user  = params.UserRepository.FindByEmail(email)
	)

	if user.IsLeft {
		either.Left(user.LeftResponse)
		return either
	}

	passwordValidation := validatePassword(params.Payload, user.RightResponse)

	if passwordValidation.IsLeft {
		either.Left(passwordValidation.LeftResponse)
		return either
	}

	userCredential := createCredentials(user.RightResponse)

	if userCredential.IsLeft {
		either.Left(user.LeftResponse)
		return either
	}

	either.Right(Response{
		User:       user.RightResponse.Map(),
		Credential: userCredential.RightResponse,
	})

	return either
}

func validatePayload(payload Payload) helpers.Either[any] {
	var either helpers.Either[any]

	err := typings.Error{
		StatusCode: 400,
		Message:    "Payload has invalid format",
	}

	if !helpers.ValidateStruct(payload, &err) {
		either.Left(err)
	} else {
		either.Right(nil)
	}

	return either
}

func validatePassword(
	payload Payload,
	user entities.UserEntity,
) helpers.Either[any] {
	var (
		either   helpers.Either[any]
		password = payload.Password
	)

	if user.PasswordHash.Validate(password) {
		either.Right(nil)
	} else {
		err := typings.Error{
			StatusCode: 401,
			Message:    "Unauthorized",
			Reason:     "Email or password is invalid",
		}

		either.Left(err)
	}

	return either
}

func createCredentials(
	user entities.UserEntity,
) helpers.Either[UserCredential] {
	var (
		either      helpers.Either[UserCredential]
		accessToken = createAccessToken(user)
	)

	if accessToken.IsLeft {
		either.Left(accessToken.LeftResponse)
		return either
	}

	refreshToken := createRefreshToken(user)

	if refreshToken.IsLeft {
		either.Left(refreshToken.LeftResponse)
		return either
	}

	either.Right(UserCredential{
		AccessToken:  accessToken.RightResponse,
		RefreshToken: refreshToken.RightResponse,
	})

	return either
}

func createAccessToken(user entities.UserEntity) helpers.Either[string] {
	return helpers.GenerateJWT(jwt.MapClaims{
		"sub": user.ID.GetValue(),
		"exp": time.Now().Add(time.Hour * 24 * 7), // 7 days
	})
}

func createRefreshToken(user entities.UserEntity) helpers.Either[string] {
	return helpers.GenerateJWT(jwt.MapClaims{
		"sub": user.ID.GetValue(),
		"exp": time.Now().Add(time.Hour * 24), // 24 hours
	})
}
