package main

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Port           string `envconfig:"PORT" default:"7194"`
	Host           string `envconfig:"HOST" default:"localhost"`
	MigrateOnStart bool   `envconfig:"MIGRATE_ON_START" default:"false"`
	SessionSalt    string `envconfig:"SESSION_SALT" required:"true"`
	DatabaseURL    string `envconfig:"DATABASE_URL" required:"true"`
}

var Config config

func InitConfig() {
	envconfig.MustProcess("", &Config)
}
