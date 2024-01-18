package server

import (
	"errors"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Env struct {
	DBHost     string `validate:"nonzero" mapstructure:"APP_DB_HOST"`
	DBPort     uint16 `validate:"nonzero" mapstructure:"APP_DB_PORT"`
	DBUser     string `validate:"nonzero" mapstructure:"APP_DB_USER"`
	DBPassword string `validate:"nonzero" mapstructure:"APP_DB_PASSWORD"`
	DBName     string `validate:"nonzero" mapstructure:"APP_DB_NAME"`

	AppPort uint16 `validate:"nonzero" mapstructure:"APP_PORT"`
}

func LoadEnv(filename string, config interface{}) error {
	for _, env := range os.Environ() {
		split := strings.Split(env, "=")

		err := viper.BindEnv(split[0])
		if err != nil {
			return err
		}
	}

	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigFile(filename)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	return viper.Unmarshal(&config)
}
