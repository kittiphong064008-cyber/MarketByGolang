package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	PostgreSQL PostgreSQL
	App        Fiber
}

type Fiber struct {
	Host string
	Port string
}

type PostgreSQL struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func Load() Configs {
	godotenv.Load()
	return Configs{
		App: Fiber{
			Host: os.Getenv("FIBER_HOST"),
			Port: os.Getenv("FIBER_PORT"),
		},
		PostgreSQL: PostgreSQL{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
	}
}
