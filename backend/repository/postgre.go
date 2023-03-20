package repository

import (
	"log"

	"github.com/McCune1224/Echo/models"
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
	DBConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	DBConnection.Logger = logger.Default.LogMode(logger.Silent)
	// DBConnection.Logger = logger.Default.LogMode(logger.Info)
	// DBConnection.Migrator().DropTable(&models.Session{})
	// DBConnection.Migrator().DropTable(&models.User{})
	// DBConnection.Migrator().DropTable(&models.Playlist{})

	DBConnection.AutoMigrate(&models.User{})
	// DBConnection.AutoMigrate(&models.Session{})
	// DBConnection.AutoMigrate(&models.Playlist{})
	// DBConnection.AutoMigrate(&models.Track{})
}
