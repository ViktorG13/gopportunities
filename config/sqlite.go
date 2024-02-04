package config

import (
	"os"

	"github.com/ViktorG13/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("SQLite")
	dbPath := "./db/main.db"

	// Check if The Database Exists.
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database File Not Found, Creating a new...")

		// Create The Folder And File Of The Database.
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Create a Database And Connect.
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("SQLite Opening Error: %v", err)
		return nil, err
	}
	// Migrate The Schema.
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("SQLite Auto Migration Error: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}
