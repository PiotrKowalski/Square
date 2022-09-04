package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	GrpcPort string `envconfig:"GRPC_PORT" default:"30000" json:"GRPC_PORT"`
	RestPort string `envconfig:"REST_PORT" default:"40000" json:"REST_PORT"`

	EchoUrl   string `envconfig:"ECHO_URL" default:"https://postman-echo.com/get"  json:"ECHO_URL"`
	ServeHTTP bool   `envconfig:"SERVE_HTTP" default:"true" required:"false" json:"SERVE_HTTP"`
}

func GetConfig() (Config, error) {
	var conf Config
	err := envconfig.Process("Service", &conf)
	if err != nil {
		return Config{}, fmt.Errorf("can't load config: %w", err)
	}

	return conf, nil
}
