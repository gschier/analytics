package main

import "github.com/kelseyhightower/envconfig"

type config struct {
	Port           string `envconfig:"PORT" default:"7194"`
	Host           string `envconfig:"HOST" default:"localhost"`
	DatabaseURL    string `envconfig:"DATABASE_URL" required:"true"`
	MigrateOnStart bool   `envconfig:"MIGRATE_ON_START" default:"false"`
}

var Config config

func init() {
	envconfig.MustProcess("", &Config)
}
