package sqlite

import (
	"comi-track/internal/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, domain.NewAppError(domain.ErrInternal, "failed to connect to database")
	}
	db.AutoMigrate(&ArticleModel{})

	return db, nil
}
