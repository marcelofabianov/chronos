package config

import (
	"encoding/json"

	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		Server        ServerConfig
		Logger        LoggerConfig
		Database      DatabaseConfig
		AuditDatabase DatabaseConfig
		Auth          AuthConfig
	}

	ServerConfig struct {
		Host      string
		Port      int
		RateLimit int
	}

	LoggerConfig struct {
		Level string
	}

	DatabaseConfig struct {
		Driver          string
		Host            string
		Port            int
		User            string
		Password        string
		Name            string
		SSLMode         string
		MaxOpenConns    int
		MaxIdleConns    int
		ConnMaxLifetime int
		ConnMaxIdleTime int
	}

	AuthConfig struct {
		Cors CorsConfig
	}

	CorsConfig struct {
		AllowedOrigins   []string
		AllowedMethods   []string
		AllowedHeaders   []string
		ExposedHeaders   []string
		AllowCredentials bool
	}
)

func LoadConfig() (*AppConfig, error) {
	v := viper.New()

	v.BindEnv("server.host", "APP_API_HOST")
	v.BindEnv("server.port", "APP_API_PORT")
	v.BindEnv("server.ratelimit", "APP_API_RATE_LIMIT")

	v.BindEnv("logger.level", "APP_LOGGER_LEVEL")

	v.BindEnv("database.driver", "APP_DB_DRIVER")
	v.BindEnv("database.host", "APP_DB_HOST")
	v.BindEnv("database.port", "APP_DB_PORT")
	v.BindEnv("database.user", "APP_DB_USER")
	v.BindEnv("database.password", "APP_DB_PASSWORD")
	v.BindEnv("database.name", "APP_DB_NAME")
	v.BindEnv("database.sslmode", "APP_DB_SSL_MODE")
	v.BindEnv("database.maxopenconns", "APP_DB_MAXOPENCONNS")
	v.BindEnv("database.maxidleconns", "APP_DB_MAXIDLECONNS")
	v.BindEnv("database.connmaxlifetime", "APP_DB_CONNMAXLIFETIME")
	v.BindEnv("database.connmaxidletime", "APP_DB_CONNMAXIDLETIME")

	v.BindEnv("auth.cors.allowedorigins", "APP_AUTH_CORS_ALLOWEDORIGINS")
	v.BindEnv("auth.cors.allowedmethods", "APP_AUTH_CORS_ALLOWEDMETHODS")
	v.BindEnv("auth.cors.allowedheaders", "APP_AUTH_CORS_ALLOWEDHEADERS")
	v.BindEnv("auth.cors.exposedheaders", "APP_AUTH_CORS_EXPOSEDHEADERS")
	v.BindEnv("auth.cors.allowcredentials", "APP_AUTH_CORS_ALLOWCREDENTIALS")

	// defaults...
	v.SetDefault("server.host", "0.0.0.0")
	v.SetDefault("server.port", 8080)
	v.SetDefault("logger.level", "info")
	v.SetDefault("server.ratelimit", 100)
	v.SetDefault("database.maxOpenConns", 10)
	v.SetDefault("database.maxIdleConns", 10)
	v.SetDefault("database.connMaxLifetime", 5)
	v.SetDefault("database.connMaxIdleTime", 5)

	var cfg AppConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *AppConfig) ToJSON() (string, error) {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
