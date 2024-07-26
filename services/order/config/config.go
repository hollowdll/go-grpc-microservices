package config

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

const (
	EnvPrefix                      string = "ORDER"
	GrpcPortConfig                 string = "GRPC_PORT"
	ApplicationModeConfig          string = "APPLICATION_MODE"
	InventoryServiceHostConfig     string = "INVENTORY_SERVICE_HOST"
	InventoryServiceGrpcPortConfig string = "INVENTORY_SERVICE_GRPC_PORT"
	PaymentServiceHostConfig       string = "PAYMENT_SERVICE_HOST"
	PaymentServiceGrpcPortConfig   string = "PAYMENT_SERVICE_GRPC_PORT"

	DefaultGrpcPort                 int    = 9002
	DefaultApplicationMode          string = "development"
	DefaultInventoryServiceHost     string = "localhost"
	DefaultInventoryServiceGrpcPort int    = 9001
	DefaultPaymentServiceHost       string = "localhost"
	DefaultPaymentServiceGrpcPort   int    = 9000
)

var defaultConfigs = map[string]interface{}{
	GrpcPortConfig:                 DefaultGrpcPort,
	ApplicationModeConfig:          DefaultApplicationMode,
	InventoryServiceHostConfig:     DefaultInventoryServiceHost,
	InventoryServiceGrpcPortConfig: DefaultInventoryServiceGrpcPort,
	PaymentServiceHostConfig:       DefaultPaymentServiceHost,
	PaymentServiceGrpcPortConfig:   DefaultPaymentServiceGrpcPort,
}

type Config struct {
	GrpcPort                 int
	ApplicationMode          string
	InventoryServiceHost     string
	InventoryServiceGrpcPort int
	PaymentServiceHost       string
	PaymentServiceGrpcPort   int
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
		GrpcPort:                 convertPort(viper.GetString(GrpcPortConfig)),
		ApplicationMode:          viper.GetString(ApplicationModeConfig),
		InventoryServiceHost:     viper.GetString(InventoryServiceHostConfig),
		InventoryServiceGrpcPort: convertPort(viper.GetString(InventoryServiceGrpcPortConfig)),
		PaymentServiceHost:       viper.GetString(PaymentServiceHostConfig),
		PaymentServiceGrpcPort:   convertPort(viper.GetString(PaymentServiceGrpcPortConfig)),
	}
}

func InitConfig() {
	viper.SetConfigName("orderservice-config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.SetDefault(GrpcPortConfig, DefaultGrpcPort)
	viper.SetDefault(ApplicationModeConfig, DefaultApplicationMode)
	viper.SetDefault(InventoryServiceHostConfig, DefaultInventoryServiceHost)
	viper.SetDefault(InventoryServiceGrpcPortConfig, DefaultInventoryServiceGrpcPort)
	viper.SetDefault(PaymentServiceHostConfig, DefaultPaymentServiceHost)
	viper.SetDefault(PaymentServiceGrpcPortConfig, DefaultPaymentServiceGrpcPort)

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
