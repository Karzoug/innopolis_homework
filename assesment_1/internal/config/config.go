package config

import (
	"assesment_1/internal/delivery/http"
	"assesment_1/internal/service"
)

type Config struct {
	HTTP    http.Config    `envPrefix:"HTTP_"`
	Service service.Config `envPrefix:"SERVICE_"`
}
