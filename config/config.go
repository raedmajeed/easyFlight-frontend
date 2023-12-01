package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Parameters struct {
	PORT        string `mapstructure:"PORT"`
	SECRETKEY   string `mapstructure:"SECRETKEY""`
	BACKENDPORT string `mapstructure:"BACKENDPORT"`
}

func Config() (*Parameters, error) {
	viper.SetConfigFile("../../.env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading from .env file %v", err.Error())
	}

	var params Parameters
	if err := viper.Unmarshal(&params); err != nil {
		return nil, fmt.Errorf("error unmarshalling from .env file %v", err.Error())
	}
	return &params, nil
}
