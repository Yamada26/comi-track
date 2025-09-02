package domain

import "comi-track/internal/common"

type Article struct {
	id    int
	title string
}

func (a *Article) GetID() int {
	return a.id
}

func (a *Article) GetTitle() string {
	return a.title
}

func NewArticle(id int, title string) (*Article, error) {
	if title == "" {
		return nil, common.NewAppError(common.ErrInvalid, "title must not be empty")
	}

	return &Article{id: id, title: title}, nil
}

type ArticleRepository interface {
	FindById(id int) (*Article, error)
	Create(article *Article) (*Article, error)
}
