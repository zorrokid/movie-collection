package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Database environement settings
type Database struct {
	User     string
	Password string
	Port     string
	Database string
	Host     string
}

var initialized = false

// InitEnv initilalized the environment variables
func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	initialized = true
}

// GetEnvDB return environment varialbes for database connection
func GetEnvDB() Database {
	if initialized == false {
		InitEnv()
	}
	db := Database{
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("DBPASSWORD"),
		Port:     os.Getenv("DBPORT"),
		Database: os.Getenv("DBNAME"),
		Host:     os.Getenv("DBHOST"),
	}
	return db
}
