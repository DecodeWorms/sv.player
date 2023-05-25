package condig

import (
	"os"
	"strconv"
)

// declare a constant for db config and server config
const (
	DatabaseName = "SOCCER_METRICS"
	DatabaseUrl  = "DATABASE_URL"
)

// Config holds fields for configuration
type Config struct {
	DatabaseName string
	DatabaseUrl  string
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
	return &Config{
		DatabaseName: databaseName,
		DatabaseUrl:  databaseUrl,
	}
}
