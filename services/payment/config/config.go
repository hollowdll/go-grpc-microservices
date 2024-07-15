package config

import (
	"log"
	"os"
	"strconv"
)

const (
	GrpcPortEnv        string = "PAYMENT_GRPC_PORT"
	ApplicationModeEnv string = "PAYMENT_APPLICATION_MODE"
)

type Config struct {
	GrpcPort        int
	ApplicationMode string
}

func NewConfig() *Config {
	return &Config{
		GrpcPort:        convertPort(getEnvValue(GrpcPortEnv)),
		ApplicationMode: getEnvValue(ApplicationModeEnv),
	}
}

func (c *Config) IsDevelopmentMode() bool {
	return c.ApplicationMode == "development"
}

func getEnvValue(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("%s environment variable is missing", key)
	}

	return value
}

func convertPort(portStr string) int {
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("port %s is invalid", portStr)
	}

	return port
}
