package internal

import (
	"backend/config"
	"os"
)

func SetupEnvironment() {
	os.Chdir("../../")
	os.Setenv("ENV", "test")
	config.SetupEnvironmentVariables()
}
