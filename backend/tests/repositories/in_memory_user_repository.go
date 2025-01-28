package repositories

import (
	"backend/internal/application/helpers"
	"backend/internal/application/typings"
	"backend/internal/domain/entities"
)

type InMemoryUserRepository struct{}

var userDB = map[string]entities.UserEntity{}

func (InMemoryUserRepository) Insert(
	user entities.UserEntity,
) helpers.Either[any] {
	id := user.ID.GetValue()

	userDB[id] = user

	return helpers.Either[any]{
		IsLeft:        false,
		IsRight:       true,
		RightResponse: nil,
	}
}

func (InMemoryUserRepository) FindByEmail(
	email string,
) helpers.Either[entities.UserEntity] {
	var (
		either helpers.Either[entities.UserEntity]
		user   entities.UserEntity
	)

	for id := range userDB {
		user = userDB[id]

		if user.Email == email {
			either.Right(user)
			return either
		}
	}

	err := typings.Error{
		StatusCode: 400,
		Message:    "Failed to find user by email",
		Reason:     "Email not found",
	}

	either.Left(err)
	return either
}

func (InMemoryUserRepository) DeleteByID(id string) helpers.Either[any] {
	var (
		either      helpers.Either[any]
		emptyResult = entities.UserEntity{}
		user        = userDB[id]
	)

	if user == emptyResult {
		err := typings.Error{
			StatusCode: 400,
			Message:    "Failed to find user by email",
			Reason:     "Email not found",
		}

		either.Left(err)
	} else {
		delete(userDB, id)
		either.Right(nil)
	}

	return either
}
