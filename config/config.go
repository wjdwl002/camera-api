package config

import "os"

func GetDSN() string {
	return os.Getenv("DB_DSN")
}
