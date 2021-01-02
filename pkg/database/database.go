package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initialized the database connectio
func InitDB() (*gorm.DB, error) {

	dsn := "host=localhost port=5432 user=movies dbname=moviesdb password=movies123"
	// Connect to postgres database
	DBCon, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect database")
	}

	return DBCon, err
}
