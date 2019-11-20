package config

import (
	"fmt"

	"github.com/anodamobi/go-tg-api/db"

	"github.com/caarlos0/env"
)

type Database struct {
	Name         string `env:"API_DATABASE_NAME,required"`
	Host         string `env:"API_DATABASE_HOST,required"`
	Port         int    `env:"API_DATABASE_PORT,required"`
	User         string `env:"API_DATABASE_USER,required"`
	Password     string `env:"API_DATABASE_PASSWORD,required"`
	SSL          string `env:"API_DATABASE_SSL,required"`
	MaxOpenConns int    `env:"API_DATABASE_MAX_OPEN_CONNS" envDefault:"100"`
	MaxIdleConns int    `env:"API_DATABASE_MAX_IDLE_CONNS" envDefault:"100"`
}

func (d Database) URL() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s", d.Host, d.Port, d.User, d.Password, d.Name, d.SSL)
}

func (c *ConfigImpl) DB() *db.DB {
	if c.db != nil {
		return c.db
	}

	c.Lock()
	defer c.Unlock()

	var database Database
	if err := env.Parse(&database); err != nil {
		panic(err)
	}

	dbInstance, err := db.New(database.URL(), database.MaxOpenConns, database.MaxIdleConns)
	if err != nil {
		panic(err)
	}

	c.db = dbInstance
	return c.db
}
