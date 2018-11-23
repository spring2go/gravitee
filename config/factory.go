package config

import "github.com/jinzhu/configor"

// DefaultConfig ...
// Let's start with some sensible defaults
var DefaultConfig = &Config{
	Database: DatabaseConfig{
		Type:         "mysql",
		Host:         "localhost",
		Port:         3306,
		User:         "gravitee_oauth2_server",
		Password:     "",
		DatabaseName: "gravitee_oauth2_server",
		MaxIdleConns: 5,
		MaxOpenConns: 5,
	},
	Oauth: OauthConfig{
		AccessTokenLifetime:  3600,    // 1 hour
		RefreshTokenLifetime: 1209600, // 14 days
		AuthCodeLifetime:     3600,    // 1 hour
	},
	Session: SessionConfig{
		Secret:   "test_secret",
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HTTPOnly: true,
	},
	IsDevelopment: true,
}

// NewDefaultConfig returns *Config struct for testing purpose
func NewDefaultConfig() *Config {
	return DefaultConfig
}

// NewConfig loads configuration from config file
func NewConfig(configFile string) *Config {
	if configFile != "" {
		config := &Config{}
		configor.Load(config, configFile)
		return config
	}

	return NewDefaultConfig()
}
