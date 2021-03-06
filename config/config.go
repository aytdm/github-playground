package config

import (
  "log"

  "github.com/spf13/viper"
)

type Config struct {
  Github struct {
    Owner string
    Repository string
    Token string
  }
}

var configInstance *Config

func LoadConfig() (config *Config, err error) {
  if configInstance != nil {
    config = configInstance
    return
  }

  viper.SetConfigName("setting")
  viper.AddConfigPath(".")
  viper.AutomaticEnv()

  if err := viper.ReadInConfig(); err != nil {
    log.Fatalf("Error while reading config file %s", err)
  }

  err = viper.Unmarshal(&config)
  if err != nil {
    log.Fatalf("unable to decode into config struct, %v", err)
  }
  configInstance = config

  return
}
