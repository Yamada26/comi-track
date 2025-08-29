package usecase

import (
	"comi-track/internal/domain"
	"comi-track/pkg/logger"
)

type ArticleUsecase struct {
	articleRepository domain.ArticleRepository
}

func NewArticleUsecase(ar domain.ArticleRepository) *ArticleUsecase {
	return &ArticleUsecase{articleRepository: ar}
}

type ArticleDTO struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func (au *ArticleUsecase) CreateArticle(article *domain.Article) (*ArticleDTO, error) {
	logger.Logger.Info("Usecase: CreateArticle called", "title", article.GetTitle())

	created, err := au.articleRepository.Create(article)
	if err != nil {
		logger.Logger.Error("Usecase: CreateArticle failed", "error", err)
		return nil, err
	}

	logger.Logger.Info("Usecase: CreateArticle succeeded", "id", created.GetId())
	return &ArticleDTO{
		ID:    created.GetId(),
		Title: created.GetTitle(),
	}, nil
}

func (au *ArticleUsecase) GetArticleById(id int) (*ArticleDTO, error) {
	logger.Logger.Info("Usecase: GetArticleById called", "id", id)

	article, err := au.articleRepository.FindById(id)
	if err != nil {
		logger.Logger.Error("Usecase: GetArticleById failed", "id", id, "error", err)
		return nil, err
	}

	logger.Logger.Info("Usecase: GetArticleById succeeded", "id", article.GetId())
	return &ArticleDTO{
		ID:    article.GetId(),
		Title: article.GetTitle(),
	}, nil
}
