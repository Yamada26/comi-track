package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func InitDB(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&ArticleModel{})

	return db, nil
}
