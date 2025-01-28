package main

import (
	"backend/config"
	"backend/internal/drivers/httpserver"
)

func main() {
	config.SetupEnvironmentVariables()
	httpserver.Run()
}
