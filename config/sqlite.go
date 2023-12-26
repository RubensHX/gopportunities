package config

import (
	"github.com/RubensHX/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	db, err := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})
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
