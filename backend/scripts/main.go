package main

import (
	"backend/config"
	"backend/scripts/migrations"
	"os"
)

func main() {
	config.SetupEnvironmentVariables()

	arg := os.Args[1]

	switch arg {
	case "migration:up":
		migrations.Up()
	case "migration:down":
		migrations.Down()
	default:
		panic("The command argument is invalid")
	}
}
