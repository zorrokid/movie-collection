package database

import (
	"fmt"

	"github.com/zorrokid/movieCollection/pkg/models"

	"github.com/zorrokid/movieCollection/pkg/database/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initialized the database connectio
func InitDB() (*gorm.DB, error) {

	db := env.GetEnvDB()

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", db.Host, db.Port, db.User, db.Database, db.Password)
	// Connect to postgres database
	DBCon, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect database")
	}

	return DBCon, err
}

// MigrateDB runs migrations to database
func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.CaseType{}, &models.Condition{}, &models.CollectionStatus{}, &models.Compilation{})
}
