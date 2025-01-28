package e2e

import (
	"backend/internal/drivers/httpserver"
	"backend/internal/infrastructure"
	"backend/internal/persistence/models"
	"backend/tests/e2e/internal"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestHandleUserCreation(t *testing.T) {
	internal.SetupEnvironment()

	var email = faker.Internet().Email()

	t.Cleanup(func() {
		var postgres = infrastructure.Postgres{}

		postgres.Connect()

		defer postgres.Disconnect()

		postgres.DB.Delete(models.User{}, "email=?", email)
	})

	template := fmt.Sprintf(`{
		"email": "%s",
		"password": "%s",
		"role": "user"
	}`, email, "abcABC123!@")

	body := bytes.NewBuffer([]byte(template))

	var (
		server = httptest.NewServer(httpserver.Setup())
		url    = server.URL + "/users"
		token  = internal.GenerateToken()
	)

	req, err := http.NewRequest("POST", url, body)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	require.Nil(t, err)

	var client = &http.Client{}

	res, err := client.Do(req)

	require.Nil(t, err)
	require.Equal(t, 201, res.StatusCode)
}
