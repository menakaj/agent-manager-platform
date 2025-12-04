// Copyright (c) 2025, WSO2 LLC (http://www.wso2.com). All Rights Reserved.
//
// This software is the property of WSO2 LLC and its suppliers, if any.
// Dissemination of any information or reproduction of any material contained
// herein is strictly forbidden, unless permitted by WSO2 in accordance with
// the WSO2 Commercial License available at http://wso2.com/licenses.
// For specific language governing the permissions and limitations under
// this license, please see the license as well as any agreement you've
// entered into with WSO2 governing the purchase of this software and any
// associated services.

package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds all configuration for the tracing service
type Config struct {
	Server     ServerConfig
	OpenSearch OpenSearchConfig
}

// ServerConfig holds HTTP server configuration
type ServerConfig struct {
	Port int
}

// OpenSearchConfig holds OpenSearch connection configuration
type OpenSearchConfig struct {
	Address  string
	Username string
	Password string
	Index    string
}

// Load loads configuration from environment variables with defaults
func Load() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Port: getEnvAsInt("TRACES_OBSERVER_PORT", 9098),
		},
		OpenSearch: OpenSearchConfig{
			Address:  getEnv("OPENSEARCH_ADDRESS", "http://localhost:9200"),
			Username: getEnv("OPENSEARCH_USERNAME", "admin"),
			Password: getEnv("OPENSEARCH_PASSWORD", "admin"),
			Index:    getEnv("OPENSEARCH_TRACE_INDEX", "custom-otel-span-index"),
		},
	}

	// Validate
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	if c.Server.Port <= 0 || c.Server.Port > 65535 {
		return fmt.Errorf("invalid server port: %d", c.Server.Port)
	}
	if c.OpenSearch.Address == "" {
		return fmt.Errorf("opensearch address is required")
	}
	return nil
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
