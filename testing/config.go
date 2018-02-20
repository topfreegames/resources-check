package testing

import (
	"strings"

	"github.com/spf13/viper"
)

// GetDefaultConfig returns the configuration at ./config/test.yaml
func GetDefaultConfig() (*viper.Viper, error) {
	cfg := viper.New()
	cfg.SetConfigName("test")
	cfg.SetConfigType("yaml")
	cfg.SetEnvPrefix("rcheck")
	cfg.AddConfigPath(".")
	cfg.AddConfigPath("./config")
	cfg.AddConfigPath("../config")
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	cfg.AutomaticEnv()

	// If a config file is found, read it in.
	if err := cfg.ReadInConfig(); err != nil {
		return nil, err
	}

	return cfg, nil
}
