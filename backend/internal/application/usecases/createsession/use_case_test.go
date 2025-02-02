package createsession

import (
	"backend/internal/application/helpers"
	"backend/internal/application/typings"
	"backend/internal/domain/entities"
	"backend/tests/repositories"
	"testing"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestCreateSession(t *testing.T) {
	t.Run("With all valid data", func(t *testing.T) {
		var (
			user           entities.UserEntity
			userRepository = repositories.InMemoryUserRepository{}
			password       = "ABCabc123!@"
		)

		user.Create(entities.UserEntityCreation{
			Email:    faker.Internet().Email(),
			Password: password,
			Role:     "user",
		})

		userRepository.Insert(user)

		result := Run(Params{
			UserRepository: userRepository,
			Payload: Payload{
				Email:    user.Email,
				Password: password,
			},
		})

		require.True(t, result.IsRight)
		require.NotEmpty(t, result.RightResponse)
		require.NotEmpty(t, result.RightResponse.User)
		require.NotEmpty(t, result.RightResponse.Credential)
	})

	t.Run("With non-existent user", func(t *testing.T) {
		var (
			userRepository = repositories.InMemoryUserRepository{}
		)

		result := Run(Params{
			UserRepository: userRepository,
			Payload: Payload{
				Email:    faker.Internet().Email(),
				Password: "ABCabc123!@",
			},
		})

		expectedResult := helpers.Either[Response]{
			IsLeft:  true,
			IsRight: false,
			LeftResponse: typings.Error{
				StatusCode: 400,
				Message:    "Failed to find user by email",
				Reason:     "Email not found",
			},
		}

		require.True(t, result.IsLeft)
		require.Equal(t, expectedResult, result)
	})

	t.Run("With invalid password", func(t *testing.T) {
		var (
			user           entities.UserEntity
			userRepository = repositories.InMemoryUserRepository{}
			password       = "ABCabc123!@"
		)

		user.Create(entities.UserEntityCreation{
			Email:    faker.Internet().Email(),
			Password: password,
			Role:     "user",
		})

		userRepository.Insert(user)

		result := Run(Params{
			UserRepository: userRepository,
			Payload: Payload{
				Email:    user.Email,
				Password: faker.Internet().Password(8, 48),
			},
		})

		expectedResult := helpers.Either[Response]{
			IsLeft:  true,
			IsRight: false,
			LeftResponse: typings.Error{
				StatusCode: 401,
				Message:    "Unauthorized",
				Reason:     "Email or password is invalid",
			},
		}

		require.Equal(t, expectedResult, result)
	})
}
