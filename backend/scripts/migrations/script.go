package migrations

import (
	"backend/internal/infrastructure"
	"backend/internal/persistence/models"
)

func Up() {
	postgres := infrastructure.Postgres{}

	postgres.Connect()

	defer postgres.Disconnect()

	err := postgres.DB.AutoMigrate(models.User{})

	if err != nil {
		panic(err)
	}
}

func Down() {
	postgres := infrastructure.Postgres{}

	postgres.Connect()

	defer postgres.Disconnect()

	postgres.DB.Exec("DROP TABLE users")
}
