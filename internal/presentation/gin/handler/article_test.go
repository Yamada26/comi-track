package handler_test

import (
	"comi-track/internal/presentation/gin/handler"
	"comi-track/internal/usecase"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/v3/assert"
)

type StubArticleUsecase struct {
	GetArticleByIdFunc func(command usecase.GetArticleByIdCommand) (*usecase.ArticleDTO, error)
	CreateArticleFunc  func(command usecase.CreateArticleCommand) (*usecase.ArticleDTO, error)
}

func (au *StubArticleUsecase) CreateArticle(command usecase.CreateArticleCommand) (*usecase.ArticleDTO, error) {
	return au.CreateArticleFunc(command)
}

func (au *StubArticleUsecase) GetArticleById(command usecase.GetArticleByIdCommand) (*usecase.ArticleDTO, error) {
	return au.GetArticleByIdFunc(command)
}

func TestGetArticleById_Success(t *testing.T) {
	mockUsecase := &StubArticleUsecase{
		GetArticleByIdFunc: func(command usecase.GetArticleByIdCommand) (*usecase.ArticleDTO, error) {
			return &usecase.ArticleDTO{
				ID:    command.ID,
				Title: "Test Article",
			}, nil
		},
	}

	handler := handler.NewArticleHandler(mockUsecase)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/articles", nil)
	ctx.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	handler.GetArticleById(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
}
