package http

type Config struct {
	Port int `env:"PORT" envDefault:"8080"`
}
