package adapters

import (
	"gopkg.in/validator.v2"
	"pilotoAPI/internal/pkg/server"
)

const (
	DEFULT_ENV = ".env"
)

var (
	serverEnv env
)

func init() {
	serverEnv = env{}
}

type env struct {
	server.Env `mapstructure:",squash"`
}

func NewEnv(filename string) error {
	if err := server.LoadEnv(filename, &serverEnv); err != nil {
		return err
	}

	if err := validator.Validate(serverEnv); err != nil {
		return err
	}

	return nil
}

func GetEnv() env {
	return serverEnv
}
