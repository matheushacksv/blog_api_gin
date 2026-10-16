package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config is the data for the app
type Config struct {
	Port        string
	DatabaseURL string
}

// Load create the config for the project
func Load() (Config, error) {
	if erro := godotenv.Load(); erro != nil {
		return Config{}, errors.New("erro to read .env file")
	}

	port := os.Getenv("PORT")
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT")

	databaseURL := createDatabaseURL(
		host,
		user,
		password,
		dbPort,
		db,
	)

	if port == "" {
		port = "8080"
	}

	return Config{
		port,
		databaseURL,
	}, nil
}

// createDatabaseURL creates the connection string for database
func createDatabaseURL(host, user, password, port, database string) string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		user, password, host, port, database,
	)
}
