package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Parameters struct {
	PORT        string `mapstructure:"PORT"`
	SECRETKEY   string `mapstructure:"SECRETKEY""`
	BACKENDPORT string `mapstructure:"BACKENDPORT"`
}

func Config() (*Parameters, error) {
	var cfg Parameters
	if err := godotenv.Load("../../.env"); err != nil {
		os.Exit(1)
	}
	cfg.PORT = os.Getenv("PORT")
	cfg.SECRETKEY = os.Getenv("SECRETKEY")
	cfg.BACKENDPORT = os.Getenv("BACKENDPORT")

	log.Println(cfg)

	return &cfg, nil
}
