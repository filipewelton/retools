package createuser

import (
	"backend/internal/application/helpers"
	"backend/internal/application/typings"
	"backend/tests/repositories"
	"testing"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestCreateUser(t *testing.T) {
	var userRepository = repositories.InMemoryUserRepository{}

	t.Run("With all valid data", func(t *testing.T) {
		email := faker.Internet().Email()

		result := Run(Params{
			UserRepository: userRepository,
			Payload: Payload{
				Email:    email,
				Password: "abcABC123!@",
				Role:     "admin",
			},
		})

		require.True(t, result.IsRight)
		require.Equal(t, email, result.RightResponse.user.Email)
	})

	t.Run("With invalid email", func(t *testing.T) {
		result := Run(Params{
			UserRepository: userRepository,
			Payload: Payload{
				Email:    faker.Internet().DomainName(),
				Password: "abcABC123!@",
				Role:     "admin",
			},
		})

		expectedResult := helpers.Either[Response]{
			IsLeft:  true,
			IsRight: false,
			LeftResponse: typings.Error{
				StatusCode: 400,
				Message:    "Payload has invalid format",
				Reason:     "Key: 'Payload.Email' Error:Field validation for 'Email' failed on the 'email' tag",
			},
		}

		require.Equal(t, expectedResult, result)
	})

	t.Run("With too short password", func(t *testing.T) {
		result := Run(Params{
			UserRepository: userRepository,
			Payload: Payload{
				Email:    faker.Internet().Email(),
				Password: "aA1!",
				Role:     "admin",
			},
		})

		expectedResult := helpers.Either[Response]{
			IsLeft:  true,
			IsRight: false,
			LeftResponse: typings.Error{
				StatusCode: 400,
				Message:    "Payload has invalid format",
				Reason:     "Key: 'Payload.Password' Error:Field validation for 'Password' failed on the 'min' tag",
			},
		}

		require.Equal(t, expectedResult, result)
	})

	t.Run("With too long password", func(t *testing.T) {
		result := Run(Params{
			UserRepository: userRepository,
			Payload: Payload{
				Email:    faker.Internet().Email(),
				Password: faker.Lorem().Characters(46) + "Ab1@",
				Role:     "admin",
			},
		})

		expectedResult := helpers.Either[Response]{
			IsLeft:  true,
			IsRight: false,
			LeftResponse: typings.Error{
				StatusCode: 400,
				Message:    "Payload has invalid format",
				Reason:     "Key: 'Payload.Password' Error:Field validation for 'Password' failed on the 'max' tag",
			},
		}

		require.Equal(t, expectedResult, result)
	})

	t.Run("With password without lowercase letters", func(t *testing.T) {
		result := Run(Params{
			UserRepository: userRepository,
			Payload: Payload{
				Email:    faker.Internet().Email(),
				Password: "ABCD1234!",
				Role:     "admin",
			},
		})

		expectedResult := helpers.Either[Response]{
			IsLeft:  true,
			IsRight: false,
			LeftResponse: typings.Error{
				StatusCode: 400,
				Message:    "Invalid password format",
				Reason:     "Password must contain lowercase letters",
			},
		}

		require.Equal(t, expectedResult, result)
	})

	t.Run("With password without uppercase letters", func(t *testing.T) {
		result := Run(Params{
			UserRepository: userRepository,
			Payload: Payload{
				Email:    faker.Internet().Email(),
				Password: "abcd1234!",
				Role:     "admin",
			},
		})

		expectedResult := helpers.Either[Response]{
			IsLeft:  true,
			IsRight: false,
			LeftResponse: typings.Error{
				StatusCode: 400,
				Message:    "Invalid password format",
				Reason:     "Password must contain uppercase letters",
			},
		}

		require.Equal(t, expectedResult, result)
	})

	t.Run("With password without numbers", func(t *testing.T) {
		result := Run(Params{
			UserRepository: userRepository,
			Payload: Payload{
				Email:    faker.Internet().Email(),
				Password: "abcdABCD!",
				Role:     "admin",
			},
		})

		expectedResult := helpers.Either[Response]{
			IsLeft:  true,
			IsRight: false,
			LeftResponse: typings.Error{
				StatusCode: 400,
				Message:    "Invalid password format",
				Reason:     "Password must contain numbers",
			},
		}

		require.Equal(t, expectedResult, result)
	})

	t.Run("With password without special characters", func(t *testing.T) {
		result := Run(Params{
			UserRepository: userRepository,
			Payload: Payload{
				Email:    faker.Internet().Email(),
				Password: "abcdABCD1",
				Role:     "admin",
			},
		})

		expectedResult := helpers.Either[Response]{
			IsLeft:  true,
			IsRight: false,
			LeftResponse: typings.Error{
				StatusCode: 400,
				Message:    "Invalid password format",
				Reason:     "Password must contain special characters",
			},
		}

		require.Equal(t, expectedResult, result)
	})

	t.Run("With empty role", func(t *testing.T) {
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
				Message:    "Payload has invalid format",
				Reason:     "Key: 'Payload.Role' Error:Field validation for 'Role' failed on the 'required' tag",
			},
		}

		require.Equal(t, expectedResult, result)
	})

	t.Run("With invalid role", func(t *testing.T) {
		result := Run(Params{
			UserRepository: userRepository,
			Payload: Payload{
				Email:    faker.Internet().Email(),
				Password: "ABCabc123!@",
				Role:     "guest",
			},
		})

		expectedResult := helpers.Either[Response]{
			IsLeft:  true,
			IsRight: false,
			LeftResponse: typings.Error{
				StatusCode: 400,
				Message:    "Payload has invalid format",
				Reason:     "Key: 'Payload.Role' Error:Field validation for 'Role' failed on the 'oneof' tag",
			},
		}

		require.Equal(t, expectedResult, result)
	})
}
