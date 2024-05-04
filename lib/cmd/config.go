package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type EnvironmentConfig struct {
	Environment        string `mapstructure:"environment"`
	ChainId            uint64 `mapstructure:"chain_id"`
	SkateWSSRPC        string `mapstructure:"skate_wss_rpc"`
	SkateHttpRPC       string `mapstructure:"skate_http_rpc"`
	SkateApp           string `mapstructure:"skate_app"`
	SkateAVS           string `mapstructure:"skate_avs"`
	HoleskyHTTPRPC     string `mapstructure:"holesky_http_rpc"`
	HoleskyWSSRPC      string `mapstructure:"holesky_wss_rpc"`
	WsETHStrategy      string `mapstructure:"wsETH_strategy"`
	EigenMetricsIPPort string `mapstructure:"eigen_metrics_ip_port_address"`
	EnableMetrics      bool   `mapstructure:"enable_metrics"`
	NodeAPIIPPort      string `mapstructure:"node_api_ip_port_address"`
	EnableNodeAPI      bool   `mapstructure:"enable_node_api"`
}

func ReadEnvironmentConfig(filename string) (*EnvironmentConfig, error) {
	// Set the name of the config file (without extension)
	viper.SetConfigName(filename)
	// Set the path to look for the config file
	viper.AddConfigPath("configs")
	// Set the config file type
	viper.SetConfigType("yaml")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
		return nil, err
	}

	// Initialize a Config struct to hold the config values
	var config EnvironmentConfig

	// Unmarshal the config into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
		return nil, err
	}
	return &config, nil
}

type SignerConfig struct {
	Address    string `mapstructure:"address"`
	Passphrase string `mapstructure:"passphrase"`
}

func ReadSignerConfig(filename string) (*SignerConfig, error) {
	// Set the name of the config file (without extension)
	viper.SetConfigName(filename)
	// Set the path to look for the config file
	viper.AddConfigPath("configs/signer")
	// Set the config file type
	viper.SetConfigType("yaml")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
		return nil, err
	}

	// Initialize a Config struct to hold the config values
	var config SignerConfig

	// Unmarshal the config into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
		return nil, err
	}
	return &config, nil
}

func ReadConfig[T any](subDir string, filename string) (*T, error) {
	// Set the name of the config file (without extension)
	viper.SetConfigName(filename)
	// Set the path to look for the config file
	viper.AddConfigPath(fmt.Sprintf("configs%s", subDir))
	// Set the config file type
	viper.SetConfigType("yaml")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
		return nil, err
	}

	// Initialize a Config struct to hold the config values
	var config T

	// Unmarshal the config into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
		return nil, err
	}
	return &config, nil
}
