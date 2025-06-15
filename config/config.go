package config

import "os"

func GetDSN() string {
	return os.Getenv("DB_DSN")
}

func GetUsername() string {
	return os.Getenv("DB_USERNAME")
}

func GetPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func GetHost() string {
	return os.Getenv("DB_HOST")
}

func GetDBPort() string {
	return os.Getenv("DB_PORT")
}

func GetDatabase() string {
	return os.Getenv("DB_DATABASE")
}