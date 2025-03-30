package app

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	AppName string `envconfig:"APP_NAME" default:"go-boilerplate"`
	AppHost  string `envconfig:"APP_HOST" required:"true"`
	AppPort string `envconfig:"APP_PORT" required:"true"`

	OtelEndpoint string `envconfig:"OTLP_ENDPOINT" required:"false"`
}

var env envConfig

func BindENV() {
	bindDotENV()

	err := envconfig.Process("", &env)
	if err != nil {
		printSpecUsage()
		panic(err)
	}
}

func ENV() envConfig {
	return env
}

func printSpecUsage() {
	err := envconfig.Usage("", &env)
	if err != nil {
		panic(err)
	}
}

func bindDotENV() {
	err := godotenv.Load()
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}
	}
}
