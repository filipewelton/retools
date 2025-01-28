package config

import (
	"os"

	"github.com/joho/godotenv"
)

type envSchema struct {
	HTTP_SERVER_ADDR string
	POSTGRES_URI     string
	JWT_SECRET       []byte
	JWT_SUBJECT      string
}

var Env envSchema

func SetupEnvironmentVariables() {
	env := os.Getenv("ENV")

	var dotEnv string

	switch env {
	case "test":
		dotEnv = ".env.test"
	case "production":
		return
	case "staging":
		return
	default:
		dotEnv = ".env.development"
	}

	err := godotenv.Load(dotEnv)

	if err != nil {
		panic(err)
	}

	Env.setValue()
}

func (e *envSchema) setValue() {
	var (
		HTTP_SERVER_ADDR = os.Getenv("HTTP_SERVER_ADDR")
		POSTGRES_URI     = os.Getenv("POSTGRES_URI")
		JWT_SECRET       = os.Getenv("JWT_SECRET")
		JWT_SUBJECT      = os.Getenv("JWT_SUBJECT")
	)

	e.HTTP_SERVER_ADDR = HTTP_SERVER_ADDR
	e.POSTGRES_URI = POSTGRES_URI
	e.JWT_SECRET = []byte(JWT_SECRET)
	e.JWT_SUBJECT = JWT_SUBJECT
}
