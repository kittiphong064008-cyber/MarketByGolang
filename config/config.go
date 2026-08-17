package configs

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Configs struct {
	App        Fiber      `mapstructure:"app"`
	PostgreSQL PostgreSQL `mapstructure:"postgres"`
}

type Fiber struct {
	Host string `mapstructure:"FIBER_HOST"`
	Port string `mapstructure:"FIBER_PORT"`
}

type PostgreSQL struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	Database string `mapstructure:"DB_DATABASE"`
	SSLMode  string `mapstructure:"DB_SSLMODE"`
}

func Load() Configs {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	setDefaults(v)
	bindEnvVars(v)

	if err := v.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if errors.As(err, &notFound) {
			log.Println("Not Found config.yml")
		} else {
			log.Fatalf("can not read config.yml")
		}
	}

	var cfg Configs
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("can not decode config")
	}
	return cfg
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("app.FIBER_HOST", "localhost")
	v.SetDefault("app.FIBER_PORT", "3000")
	v.SetDefault("postgres.DB_HOST", "localhost")
	v.SetDefault("postgres.DB_PORT", "5432")
	v.SetDefault("postgres.DB_SSLMODE", "disable")
}

func bindEnvVars(v *viper.Viper) {
	v.AutomaticEnv()

	_ = v.BindEnv("app.FIBER_HOST", "FIBER_HOST")
	_ = v.BindEnv("app.FIBER_PORT", "FIBER_PORT")

	_ = v.BindEnv("postgres.DB_HOST", "DB_HOST")
	_ = v.BindEnv("postgres.DB_PORT", "DB_PORT")
	_ = v.BindEnv("postgres.DB_USERNAME", "DB_USERNAME")
	_ = v.BindEnv("postgres.DB_PASSWORD", "DB_PASSWORD")
	_ = v.BindEnv("postgres.DB_DATABASE", "DB_DATABASE")
	_ = v.BindEnv("postgres.DB_SSLMODE", "DB_SSLMODE")
}
