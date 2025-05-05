package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	OllamaURL string
	APIURL    string
	MCPURL    string
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{
		OllamaURL: viper.GetString("ollama_url"),
		APIURL:    viper.GetString("api_url"),
		MCPURL:    viper.GetString("mcp_url"),
	}
	return cfg, nil
}
