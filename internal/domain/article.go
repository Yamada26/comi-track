package domain

type Article struct {
	id    int
	title string
}

func (a *Article) GetId() int {
	return a.id
}

func (a *Article) GetTitle() string {
	return a.title
}

func NewArticle(id int, title string) (*Article, error) {
	if title == "" {
		return nil, NewAppError(ErrInvalid, "title must not be empty", nil)
	}

	return &Article{id: id, title: title}, nil
}

type ArticleRepository interface {
	FindById(id int) (*Article, error)
	Create(article *Article) (*Article, error)
}
