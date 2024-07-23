package config

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

const (
	EnvPrefix                  string = "ORDER"
	GrpcPortConfig             string = "GRPC_PORT"
	ApplicationModeConfig      string = "APPLICATION_MODE"
	InventoryServiceHostConfig string = "INVENTORY_SERVICE_HOST"
	InventoryServicePortConfig string = "INVENTORY_SERVICE_PORT"
	PaymentServiceHostConfig   string = "PAYMENT_SERVICE_HOST"
	PaymentServicePortConfig   string = "PAYMENT_SERVICE_PORT"

	DefaultGrpcPort             int    = 9002
	DefaultApplicationMode      string = "development"
	DefaultInventoryServiceHost string = "localhost"
	DefaultInventoryServicePort int    = 9001
	DefaultPaymentServiceHost   string = "localhost"
	DefaultPaymentServicePort   int    = 9000
)

var defaultConfigs = map[string]interface{}{
	GrpcPortConfig:             DefaultGrpcPort,
	ApplicationModeConfig:      DefaultApplicationMode,
	InventoryServiceHostConfig: DefaultInventoryServiceHost,
	InventoryServicePortConfig: DefaultInventoryServicePort,
	PaymentServiceHostConfig:   DefaultPaymentServiceHost,
	PaymentServicePortConfig:   DefaultPaymentServicePort,
}

type Config struct {
	GrpcPort             int
	ApplicationMode      string
	InventoryServiceHost string
	InventoryServicePort int
	PaymentServiceHost   string
	PaymentServicePort   int
}

func NewConfig() *Config {
	log.Printf("loading configurations ...")
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
		GrpcPort:             convertPort(viper.GetString(GrpcPortConfig)),
		ApplicationMode:      viper.GetString(ApplicationModeConfig),
		InventoryServiceHost: viper.GetString(InventoryServiceHostConfig),
		InventoryServicePort: convertPort(viper.GetString(InventoryServicePortConfig)),
		PaymentServiceHost:   viper.GetString(PaymentServiceHostConfig),
		PaymentServicePort:   convertPort(viper.GetString(PaymentServicePortConfig)),
	}
}

func InitConfig() {
	viper.SetConfigName("orderservice-config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.SetDefault(GrpcPortConfig, DefaultGrpcPort)
	viper.SetDefault(ApplicationModeConfig, DefaultApplicationMode)
	viper.SetDefault(InventoryServiceHostConfig, DefaultInventoryServiceHost)
	viper.SetDefault(InventoryServicePortConfig, DefaultInventoryServicePort)
	viper.SetDefault(PaymentServiceHostConfig, DefaultPaymentServiceHost)
	viper.SetDefault(PaymentServicePortConfig, DefaultPaymentServicePort)

	viper.SetEnvPrefix(EnvPrefix)
	viper.AutomaticEnv()
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
