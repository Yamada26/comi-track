package usecase_test

import (
	"comi-track/internal/domain"
	"comi-track/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubArticleRepository struct {
	FindByIdFunc func(id int) (*domain.Article, error)
	CreateFunc   func(article *domain.Article) (*domain.Article, error)
}

func (ar *StubArticleRepository) FindById(id int) (*domain.Article, error) {
	return ar.FindByIdFunc(id)
}

func (ar *StubArticleRepository) Create(article *domain.Article) (*domain.Article, error) {
	return ar.CreateFunc(article)
}

func TestGetArticleById_Success(t *testing.T) {
	mockRepo := &StubArticleRepository{
		FindByIdFunc: func(id int) (*domain.Article, error) {
			article, err := domain.NewArticle(id, "Test article")

			if err != nil {
				return nil, err
			}

			return article, nil
		},
	}
	articleUsecase := usecase.NewArticleUsecase(mockRepo)

	articleId := 1
	command := usecase.GetArticleByIdCommand{ID: articleId}
	article, err := articleUsecase.GetArticleById(command)
	assert.NoError(t, err)
	assert.NotNil(t, article)
}
