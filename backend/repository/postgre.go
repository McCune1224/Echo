package repository

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Global variable for the database connection that can be accessed by other packages
// I'm not sure if this is the best way to do this, but it works for now until I learn more about Go
var DBConnection *gorm.DB

// Initialize the database connection
func InitDB(dsn string) {
	var err error
	DBConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	DBConnection.Logger = logger.Default.LogMode(logger.Info)
	// DBConnection.AutoMigrate(&models.User{})
	// DBConnection.AutoMigrate(&models.UserSession{})
}
