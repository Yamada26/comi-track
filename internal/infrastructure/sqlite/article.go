package sqlite

import (
	"comi-track/internal/domain"
	"comi-track/pkg/logger"

	"gorm.io/gorm"
)

type ArticleModel struct {
	ID    int    `gorm:"primaryKey;column:id"`
	Title string `gorm:"column:title"`
}

func (ArticleModel) TableName() string {
	return "articles"
}

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

/*
指定された Article を保存する。
保存に成功した場合は、保存された Article を返す。
*/
func (ar *ArticleRepository) Create(article *domain.Article) (*domain.Article, error) {
	logger.Logger.Info("Repository: Create called", "title", article.GetTitle())

	model := ArticleModel{
		ID:    article.GetId(),
		Title: article.GetTitle(),
	}

	if err := ar.db.Create(&model).Error; err != nil {
		logger.Logger.Error("Repository: failed to insert article", "error", err)
		return nil, err
	}

	logger.Logger.Info("Repository: article inserted successfully", "id", model.ID)

	return domain.NewArticle(model.ID, model.Title)
}

/*
指定された ID の Article を取得する。
存在しない場合は、gorm.ErrRecordNotFound を返す。
*/
func (ar *ArticleRepository) FindById(id int) (*domain.Article, error) {
	logger.Logger.Info("Repository: FindById called", "id", id)

	var model ArticleModel
	if err := ar.db.First(&model, id).Error; err != nil {
		logger.Logger.Error("Repository: failed to fetch article", "id", id, "error", err)
		return nil, err
	}

	logger.Logger.Info("Repository: article fetched successfully", "id", model.ID)

	return domain.NewArticle(model.ID, model.Title)
}
