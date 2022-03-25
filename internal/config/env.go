package config

import (
	"log"

	"github.com/netflix/go-env"
)

var configs Environment

const Version = "1.0.0"
const ServiceName = "go-todolist"

type Environment struct {
	Env        string `env:"ENVIRONMENT,default=dev"`
	Port       string `env:"PORT,default=8000"`
	Debug      bool   `env:"DEBUG,default=false"`
	DBHost     string `env:"DB_HOST,default=localhost"`
	DBPort     string `env:"DB_Port,default=5432"`
	DBName     string `env:"DB_NAME,default=gotask"`
	DBUser     string `env:"DB_USER,default=postgres"`
	DBPassword string `env:"DB_PASSWORD,default=secret"`
	SecretKey  string `env:"SECRET_KEY,default=MySuperSecretKey"`
	TraceHost  string `env:"TRACE_HOST,default=localhost:14268"`
}

func init() {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
	}
	configs = environment
}

func GetEnv() Environment {
	return configs
}
