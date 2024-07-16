package config

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

const (
	EnvPrefix             string = "PAYMENT"
	GrpcPortConfig        string = "GRPC_PORT"
	ApplicationModeConfig string = "APPLICATION_MODE"

	DefaultGrpcPort        int    = 9000
	DefaultApplicationMode string = "development"
)

type Config struct {
	GrpcPort        int
	ApplicationMode string
}

func NewConfig() *Config {
	log.Printf("loading configurations ...")

	viper.SetConfigName("paymentservice")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.SetDefault(GrpcPortConfig, DefaultGrpcPort)
	viper.SetDefault(ApplicationModeConfig, DefaultApplicationMode)

	viper.SetEnvPrefix(EnvPrefix)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("failed to load configs from config file: %v", err)
	}

	return &Config{
		GrpcPort:        convertPort(viper.GetString(GrpcPortConfig)),
		ApplicationMode: viper.GetString(ApplicationModeConfig),
	}
}

func (c *Config) IsDevelopmentMode() bool {
	return c.ApplicationMode == "development"
}

func convertPort(portStr string) int {
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("port %s is invalid", portStr)
	}

	return port
}
