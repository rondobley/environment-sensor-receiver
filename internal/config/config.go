package config

import (
	"environment-sensor-receiver/internal/checkerror"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

var Config Configuration

func LoadConfig(env string) {
	config := Configuration{}
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	if env == "local" {
		viper.AddConfigPath(".")
	} else if env == "prod" {
		viper.AddConfigPath("/usr/local/etc/environment-sensor-receiver")
	} else {
		log.Error().Msgf("Unsupported environment: %s", env)
		os.Exit(1)
	}
	pwd, _ := os.Getwd()
	log.Info().Msgf("Current env: %s", env)
	log.Info().Msgf("Current working dir: %s", pwd)
	log.Info().Msgf("Using config file: %s", viper.ConfigFileUsed())
	err := viper.ReadInConfig()
	checkerror.CheckError(err)
	err = viper.Unmarshal(&config)
	checkerror.CheckError(err)

	Config = config
}
