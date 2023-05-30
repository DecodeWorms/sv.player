package config

import (
	"os"
	"strconv"
)

// declare a constant for db config and server config
const (
	DatabaseName         = "SOCCER_METRICS"
	DatabaseUrl          = "DATABASE_URL"
	DatabaseHost         = "DATABASE_HOST"
	PostgresDatabaseName = "DATABASE_NAME"
	DatabaseUserName     = "DATABASE_USER_NAME"
	DatabasePort         = "DATABASE_PORT"
)

// Config holds fields for configuration
type Config struct {
	DatabaseName         string
	DatabaseUrl          string
	DatabaseHost         string
	PostgresDatabaseName string
	DatabaseUserName     string
	DatabasePort         string
}

func (c Config) GetEnv(key, fallback string) string {
	res, exist := os.LookupEnv(key)
	if !exist {
		return fallback
	}
	return res
}

func (c Config) GetEnvBool(key string, fallback bool) bool {
	res := c.GetEnv(key, "")
	if len(res) == 0 {
		return fallback
	}
	v, err := strconv.ParseBool(res)
	if err != nil {
		return fallback
	}
	return v
}

func (c Config) GetEnvInt(key string, fallback int) int {
	res, exist := os.LookupEnv(key)
	if !exist {
		return fallback
	}
	v, err := strconv.Atoi(res)
	if err != nil {
		return fallback
	}
	return v
}

func ImportConfig(c Config) *Config {
	databaseName := c.GetEnv(DatabaseName, "soccermetrics")
	databaseUrl := c.GetEnv(DatabaseUrl, "mongodb://127.0.0.1:27017")
	databaseHost := c.GetEnv(DatabaseHost, "localhost")
	postgresDatabaseName := c.GetEnv(DatabaseName, "soccermetrics")
	databasePort := c.GetEnv(DatabasePort, "5432")
	databaseUserName := c.GetEnv(DatabaseUserName, "abdulhmeed")
	return &Config{
		DatabaseName:         databaseName,
		DatabaseUrl:          databaseUrl,
		DatabaseHost:         databaseHost,
		PostgresDatabaseName: postgresDatabaseName,
		DatabaseUserName:     databaseUserName,
		DatabasePort:         databasePort,
	}
}
