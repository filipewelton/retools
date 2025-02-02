package e2e

import (
	"backend/internal/domain/entities"
	"backend/internal/drivers/httpserver"
	"backend/internal/infrastructure"
	"backend/internal/persistence/models"
	"backend/internal/persistence/repositories"
	"backend/tests/e2e/internal"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func cleanup(email string) func() {
	return func() {
		var postgres = infrastructure.Postgres{}

		postgres.Connect()

		defer postgres.Disconnect()

		postgres.DB.Delete(models.User{}, "email=?", email)
	}
}

func TestUserRouter(t *testing.T) {
	internal.SetupEnvironment()

	t.Run("Test user creation", func(t *testing.T) {
		var email = faker.Internet().Email()

		t.Cleanup(cleanup(email))

		bodyData := fmt.Sprintf(`{
		"email": "%s",
		"password": "%s",
		"role": "user"
		}`, email, "abcABC123!@")

		body := bytes.NewBuffer([]byte(bodyData))

		var (
			server = httptest.NewServer(httpserver.Setup())
			url    = server.URL + "/users"
			token  = internal.GenerateAPICredential()
		)

		req, err := http.NewRequest("POST", url, body)

		require.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		var client = &http.Client{}

		res, err := client.Do(req)

		require.Nil(t, err)
		require.Equal(t, 201, res.StatusCode)
	})

	t.Run("Test session creation", func(t *testing.T) {
		var (
			userRepository = repositories.PostgresUserRepository{}
			user           entities.UserEntity
			password       = "ABCabc123!@"
		)

		user.Create(entities.UserEntityCreation{
			Email:    faker.Internet().Email(),
			Password: password,
			Role:     "user",
		})

		userRepository.Insert(user)

		t.Cleanup(cleanup(user.Email))

		bodyData := fmt.Sprintf(`{
		"email": "%s",
		"password": "%s"
		}`, user.Email, password)

		body := bytes.NewBuffer([]byte(bodyData))

		var (
			server = httptest.NewServer(httpserver.Setup())
			url    = server.URL + "/users/session"
			token  = internal.GenerateAPICredential()
		)

		req, err := http.NewRequest("POST", url, body)

		require.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		var client = &http.Client{}

		res, err := client.Do(req)

		require.Nil(t, err)
		require.Equal(t, 200, res.StatusCode)
	})
}
