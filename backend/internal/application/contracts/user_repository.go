package contracts

import (
	"backend/internal/application/helpers"
	"backend/internal/domain/entities"
)

type UserRepository interface {
	Insert(entities.UserEntity) helpers.Either[any]
	FindByEmail(string) helpers.Either[entities.UserEntity]
	DeleteByID(string) helpers.Either[any]
}
