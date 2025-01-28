package infrastructure

import (
	"backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Postgres struct {
	DB *gorm.DB
}

func (p *Postgres) Connect() {
	uri := config.Env.POSTGRES_URI

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Discard,
	})

	if err != nil {
		panic(err)
	}

	p.DB = db
}

func (p *Postgres) Disconnect() {
	connection, err := p.DB.DB()

	if err != nil {
		panic(err)
	}

	connection.Close()
}
