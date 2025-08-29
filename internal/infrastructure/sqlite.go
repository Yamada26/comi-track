package infrastructure

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open sqlite db: %v", err)
	}

	db.AutoMigrate(&ArticleModel{})
}

func GetDB() *gorm.DB {
	return db
}
