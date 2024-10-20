package config

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

const (
	EnvPrefix string = "INVENTORY"

	GrpcPortConfig        string = "GRPC_PORT"
	ApplicationModeConfig string = "APPLICATION_MODE"
	DBHostConfig          string = "DB_HOST"
	DBUserConfig          string = "DB_USER"
	DBPasswordConfig      string = "DB_PASSWORD"
	DBNameConfig          string = "DB_NAME"
	DBPortConfig          string = "DB_PORT"
	DBSSLModeConfig       string = "DB_SSL_MODE"

	DefaultGrpcPort        int    = 9001
	DefaultApplicationMode string = "development"
	DefaultDBHost          string = "localhost"
	DefaultDBUser          string = "service"
	DefaultDBPassword      string = "service_psw"
	DefaultDBName          string = "inventory_db"
	DefaultDBPort          int    = 5432
	DefaultDBSSLMode       string = "disable"
)

var defaultConfigs = map[string]interface{}{
	GrpcPortConfig:        DefaultGrpcPort,
	ApplicationModeConfig: DefaultApplicationMode,
	DBHostConfig:          DefaultDBHost,
	DBUserConfig:          DefaultDBUser,
	DBPasswordConfig:      DefaultDBPassword,
	DBNameConfig:          DefaultDBName,
	DBPortConfig:          DefaultDBPort,
	DBSSLModeConfig:       DefaultDBSSLMode,
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
	SSLMode  string
}

type Config struct {
	GrpcPort        int
	ApplicationMode string
	DB              DBConfig
}

func InitConfig() {
	viper.SetConfigName("inventoryservice-config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	setConfigDefaults()
	viper.SetEnvPrefix(EnvPrefix)
	viper.AutomaticEnv()
}

func LoadConfig() *Config {
	log.Printf("loading configurations ...")
	readConfigFile()
	checkConfigDefaults()

	return &Config{
		GrpcPort:        convertPort(viper.GetString(GrpcPortConfig)),
		ApplicationMode: viper.GetString(ApplicationModeConfig),
		DB: DBConfig{
			Host:     viper.GetString(DBHostConfig),
			User:     viper.GetString(DBUserConfig),
			Password: viper.GetString(DBPasswordConfig),
			DBName:   viper.GetString(DBNameConfig),
			Port:     convertPort(viper.GetString(DBPortConfig)),
			SSLMode:  viper.GetString(DBSSLModeConfig),
		},
	}
}

func readConfigFile() {
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("failed to load configs from config file: %v", err)
	}
}

func setConfigDefaults() {
	viper.SetDefault(GrpcPortConfig, DefaultGrpcPort)
	viper.SetDefault(ApplicationModeConfig, DefaultApplicationMode)
	viper.SetDefault(DBHostConfig, DefaultDBHost)
	viper.SetDefault(DBUserConfig, DefaultDBUser)
	viper.SetDefault(DBPasswordConfig, DefaultDBPassword)
	viper.SetDefault(DBNameConfig, DefaultDBName)
	viper.SetDefault(DBPortConfig, DefaultDBPort)
	viper.SetDefault(DBSSLModeConfig, DefaultDBSSLMode)
}

func checkConfigDefaults() {
	for key := range defaultConfigs {
		if isDefaultConfig(key) {
			log.Printf("using default value for config %s", key)
		} else {
			log.Printf("overwriting default value for config %s", key)
		}
	}
}

func (c *Config) IsDevelopmentMode() bool {
	return c.ApplicationMode == "development"
}

func (c *Config) IsTestingMode() bool {
	return c.ApplicationMode == "testing"
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
