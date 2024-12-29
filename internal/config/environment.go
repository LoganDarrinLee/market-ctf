package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type EnvironmentConfig struct {
	DBHost     string `envconfig:"DBHOST" required:"true"`
	DBPort     string `envconfig:"DBPORT" required:"true"`
	DBUser     string `envconfig:"DBUSER" required:"true"`
	DBPassword string `envconfig:"DBPASSWORD" required:"true"`
	DBName     string `envconfig:"DBNAME" required:"true"`
	ServerPort string `envconfig:"SERVER_PORT" required:"true"`
}

func NewEnv() *EnvironmentConfig {
	var env EnvironmentConfig
	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatal(err)
	}
	return &env
}
