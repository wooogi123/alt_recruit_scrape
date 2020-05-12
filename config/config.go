package config

import "github.com/sakirsensoy/genv"

type appConfig struct {
	DbHost     string
	DbUser     string
	DbPassword string
}

var App = &appConfig{
	DbHost:     genv.Key("DB_HOST").String(),
	DbUser:     genv.Key("DB_USER").String(),
	DbPassword: genv.Key("DB_PASSWORD").String(),
}
