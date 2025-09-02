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

type CreateArticleCommand struct {
	Title string `json:"title"`
}

func (au *ArticleUsecase) CreateArticle(command CreateArticleCommand) (*ArticleDTO, error) {
	logger.Logger.Info("Usecase: CreateArticle called", "title", command.Title)

	articleToCreate, err := domain.NewArticle(0, command.Title)
	if err != nil {
		return nil, err
	}

	created, err := au.articleRepository.Create(articleToCreate)
	if err != nil {
		return nil, err
	}

	logger.Logger.Info("Usecase: CreateArticle succeeded", "id", created.GetID())
	return &ArticleDTO{
		ID:    created.GetID(),
		Title: created.GetTitle(),
	}, nil
}

type GetArticleByIdCommand struct {
	ID int `json:"id"`
}

func (au *ArticleUsecase) GetArticleById(command GetArticleByIdCommand) (*ArticleDTO, error) {
	logger.Logger.Info("Usecase: GetArticleById called", "id", command.ID)

	article, err := au.articleRepository.FindById(command.ID)
	if err != nil {
		return nil, err
	}

	logger.Logger.Info("Usecase: GetArticleById succeeded", "id", article.GetID())
	return &ArticleDTO{
		ID:    article.GetID(),
		Title: article.GetTitle(),
	}, nil
}
