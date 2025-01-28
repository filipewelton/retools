package repositories

import (
	"backend/internal/application/helpers"
	"backend/internal/application/typings"
	"backend/internal/domain/entities"
	"backend/internal/infrastructure"
	"backend/internal/persistence/models"
	"errors"

	"gorm.io/gorm"
)

type PostgresUserRepository struct{}

func (PostgresUserRepository) Insert(
	user entities.UserEntity,
) helpers.Either[any] {
	var (
		either   helpers.Either[any]
		postgres infrastructure.Postgres
		model    models.User
	)

	postgres.Connect()

	defer postgres.Disconnect()

	model.MapToModel(user)

	if db := postgres.DB.Create(model); db.Error != nil {
		err := typings.Error{
			StatusCode: 500,
			Message:    "Failed to insert user into database",
			Reason:     db.Error.Error(),
		}

		either.Left(err)
	} else {
		either.Right(nil)
	}

	return either
}

func (PostgresUserRepository) FindByEmail(
	email string,
) helpers.Either[entities.UserEntity] {
	var (
		either   helpers.Either[entities.UserEntity]
		postgres infrastructure.Postgres
		model    models.User
	)

	postgres.Connect()

	defer postgres.Disconnect()

	db := postgres.DB.First(model, "email=?", email)

	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		err := typings.Error{
			StatusCode: 400,
			Message:    "Failed to find user by email",
			Reason:     "Email not found",
		}

		either.Left(err)
	} else if db.Error != nil {
		err := typings.Error{
			StatusCode: 500,
			Message:    "Failed to find user by email",
			Reason:     db.Error.Error(),
		}

		either.Left(err)
	} else {
		either.Right(model.MapToEntity())
	}

	return either
}

func (PostgresUserRepository) DeleteByID(id string) helpers.Either[any] {
	var (
		either   helpers.Either[any]
		postgres infrastructure.Postgres
		model    models.User
	)

	postgres.Connect()

	defer postgres.Disconnect()

	db := postgres.DB.Delete(model, "id=?", id)

	if db.Error != nil {
		err := typings.Error{
			StatusCode: 500,
			Message:    "Failed to delete user",
			Reason:     db.Error.Error(),
		}

		either.Left(err)
	} else {
		either.Right(nil)
	}

	return either
}
