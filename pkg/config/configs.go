package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type (
	defValue interface{}
	Configs  struct {
		DBUsername  string `mapstructure:"db_user"`
		DBPassword  string `mapstructure:"db_password"`
		DBName      string `mapstructure:"db_name"`
		DBHost      string `mapstructure:"db_host"`
		DBPort      uint16 `mapstructure:"db_port"`
		SSLMode     string `mapstructure:"ssl_mode"`     // postgres ssl_mode
		HttpAddress string `mapstructure:"http_address"` // address for starting grpc server
		LogLevel    string `mapstructure:"log_level"`    // log level one of: debug, info, warn, error, fatal
	}
)

func GetConfigs() Configs {
	var c Configs

	var envs = map[string]defValue{
		"db_user":      "",
		"db_password":  "",
		"db_name":      "postgres",
		"db_host":      "localhost",
		"db_port":      5432,
		"ssl_mode":     "disable",
		"http_address": ":9090",
		"log_level":    "debug",
	}

	err := initEnvs(envs)
	if err != nil {
		panic(fmt.Sprintf("Unable to init env params: %v\n", err))
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		panic(fmt.Sprintf("Unable to decode into struct, %v\n", err))
	}

	return c
}

func initEnvs(e map[string]defValue) error {
	for k, v := range e {
		viper.SetDefault(k, v)
		err := viper.BindEnv(k)
		if err != nil {
			return err
		}
	}
	return nil
}
