package main

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Port           string `envconfig:"PORT" default:"7194"`
	Host           string `envconfig:"HOST" default:"localhost"`
	DatabaseURL    string `envconfig:"DATABASE_URL" required:"true"`
	MigrateOnStart bool   `envconfig:"MIGRATE_ON_START" default:"false"`
	CacheTemplates bool   `envconfig:"CACHE_TEMPLATES" default:"false"`
	SessionSalt    string `envconfig:"SESSION_SALT" required:"true"`
}

var Config config

func InitConfig() {
	envconfig.MustProcess("", &Config)
	// b, _ := json.MarshalIndent(&Config, "|  ", "  ")
	// fmt.Printf("\n\n|\n|  %s\n|\n", b)
}
