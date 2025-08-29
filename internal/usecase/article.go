package usecase

import "comi-track/internal/domain"

type ArticleUsecase struct {
	articleRepository domain.ArticleRepository
}

func NewArticleUsecase(ar domain.ArticleRepository) *ArticleUsecase {
	return &ArticleUsecase{articleRepository: ar}
}

func (au *ArticleUsecase) CreateArticle(article *domain.Article) (*domain.Article, error) {
	return au.articleRepository.Create(article)
}

func (au *ArticleUsecase) GetArticleById(id int) (*domain.Article, error) {
	return au.articleRepository.FindById(id)
}
