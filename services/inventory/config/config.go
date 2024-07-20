package config

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

const (
	EnvPrefix             string = "INVENTORY"
	GrpcPortConfig        string = "GRPC_PORT"
	ApplicationModeConfig string = "APPLICATION_MODE"

	DefaultGrpcPort        int    = 9001
	DefaultApplicationMode string = "development"
)

var defaultConfigs = map[string]interface{}{
	GrpcPortConfig:        DefaultGrpcPort,
	ApplicationModeConfig: DefaultApplicationMode,
}

type Config struct {
	GrpcPort        int
	ApplicationMode string
}

func NewConfig() *Config {
	log.Printf("loading configurations ...")

	viper.SetConfigName("inventoryservice-config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.SetDefault(GrpcPortConfig, DefaultGrpcPort)
	viper.SetDefault(ApplicationModeConfig, DefaultApplicationMode)

	viper.SetEnvPrefix(EnvPrefix)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("failed to load configs from config file: %v", err)
	}

	for key := range defaultConfigs {
		if isDefaultConfig(key) {
			log.Printf("using default value for config %s", key)
		} else {
			log.Printf("overwriting default value for config %s", key)
		}
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

func isDefaultConfig(key string) bool {
	defaultValue, ok := defaultConfigs[key]
	if !ok {
		return false
	}
	return viper.Get(key) == defaultValue
}
