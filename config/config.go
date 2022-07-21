package config

import (
	"os"
)

type GoArbitrageConfig struct {
    PriceProxyContractAddress string
    InfuraUrl   string
}

type Config struct {
    GoArbitrageConfig GoArbitrageConfig
}

// New returns a new Config struct
func New() *Config {
    return &Config{
        GoArbitrageConfig: GoArbitrageConfig{
		PriceProxyContractAddress: getEnv("PRICE_PROXY_ADDRESS", "0x416355755f32b2710ce38725ed0fa102ce7d07e6"),
	    InfuraUrl: getEnv("INFURA_PROJECT_ID", ""),
	},
    }
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
	return value
    }

    return defaultVal
}