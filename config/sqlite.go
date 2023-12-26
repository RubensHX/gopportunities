package config

import (
	"os"

	"github.com/RubensHX/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		_, err = os.Create(dbPath)
		if err != nil {
			logger.ErrorF("Error creating sqlite: %v", err)
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
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.ErrorF("Error initializing sqlite: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opnening{})
	if err != nil {
		logger.ErrorF("Error migrating sqlite: %v", err)
		return nil, err
	}
	return db, nil
}
