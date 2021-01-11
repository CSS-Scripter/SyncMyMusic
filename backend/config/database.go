package config

import "os"

type database struct {
	URL      string
	Port     string
	User     string
	Password string
	Database string
}

// Database contains all database configuration variables
var Database database

func init() {
	Database = database{
		URL:      os.Getenv("POSTGRES_URL"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("DATABASE_NAME"),
	}
}
