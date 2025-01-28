package createuser

import (
	"backend/internal/application/contracts"
	"backend/internal/application/helpers"
	"backend/internal/application/regex"
	"backend/internal/application/typings"
	"backend/internal/domain/entities"
)

type Payload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=48"`
	Role     string `json:"role" validate:"required,oneof=admin user"`
}

type Params struct {
	UserRepository contracts.UserRepository
	Payload        Payload
}

type Response struct {
	user entities.MappedUserEntity
}

func Run(params Params) helpers.Either[Response] {
	var either helpers.Either[Response]

	validations := handleValidations(params.Payload)

	if validations.IsLeft {
		either.Left(validations.LeftResponse)
		return either
	}

	creationResult := handleCreation(params)

	if creationResult.IsLeft {
		either.Left(creationResult.LeftResponse)
		return either
	}

	user := creationResult.RightResponse

	response := Response{
		user: user.Map(),
	}

	either.Right(response)
	return either
}

func handleValidations(payload Payload) helpers.Either[any] {
	if v := validatePayload(payload); v.IsLeft {
		return v
	}

	password := payload.Password

	return validatePasswordFormat(password)
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

func validatePasswordFormat(password string) helpers.Either[any] {
	var either helpers.Either[any]

	var err = typings.Error{
		StatusCode: 400,
		Message:    "Invalid password format",
	}

	if !regex.Password.HasLowercase.MatchString(password) {
		err.Reason = "Password must contain lowercase letters"
		either.Left(err)
	} else if !regex.Password.HasUppercase.MatchString(password) {
		err.Reason = "Password must contain uppercase letters"
		either.Left(err)
	} else if !regex.Password.HasNumbers.MatchString(password) {
		err.Reason = "Password must contain numbers"
		either.Left(err)
	} else if !regex.Password.HasSpecialCharacters.MatchString(password) {
		err.Reason = "Password must contain special characters"
		either.Left(err)
	} else {
		either.Right(nil)
	}

	return either
}

func handleCreation(params Params) helpers.Either[entities.UserEntity] {
	var (
		either         helpers.Either[entities.UserEntity]
		userRepository = params.UserRepository
		email          = params.Payload.Email
	)

	emailIsAvailable := checkIfTheEmailIsAvailable(userRepository, email)

	if emailIsAvailable.IsLeft {
		either.Left(emailIsAvailable.LeftResponse)
		return either
	}

	var (
		user            = createEntity(params.Payload)
		insertionResult = userRepository.Insert(user)
	)

	if insertionResult.IsLeft {
		either.Left(insertionResult.LeftResponse)
	} else {
		either.Right(user)
	}

	return either
}

func checkIfTheEmailIsAvailable(
	userRepository contracts.UserRepository,
	email string,
) helpers.Either[any] {
	var (
		either helpers.Either[any]
		result = userRepository.FindByEmail(email)
	)

	if result.IsRight {
		err := typings.Error{
			StatusCode: 409,
			Message:    "Conflict",
			Reason:     "Email is not available",
		}

		either.Left(err)
	} else {
		either.Right(nil)
	}

	return either
}

func createEntity(payload Payload) entities.UserEntity {
	var user entities.UserEntity

	user.Create(entities.UserEntityCreation{
		Email:    payload.Email,
		Password: payload.Password,
	})

	return user
}
