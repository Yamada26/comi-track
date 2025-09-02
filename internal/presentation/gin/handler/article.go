package handler

import (
	"comi-track/internal/common"
	"comi-track/internal/usecase"
	"comi-track/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleUsecase interface {
	GetArticleById(command usecase.GetArticleByIdCommand) (*usecase.ArticleDTO, error)
	CreateArticle(command usecase.CreateArticleCommand) (*usecase.ArticleDTO, error)
}

type ArticleHandler struct {
	articleUsecase ArticleUsecase
}

func NewArticleHandler(au ArticleUsecase) *ArticleHandler {
	return &ArticleHandler{articleUsecase: au}
}

func (ah *ArticleHandler) CreateArticle(ctx *gin.Context) {
	var reqBody struct {
		Title string `json:"title"`
	}

	// リクエスト受信ログ
	logger.Logger.Info("Handler: CreateArticle called")

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		logger.Logger.Warn("Handler: invalid request body", "error", err)
		ctx.Error(err)
		return
	}

	command := usecase.CreateArticleCommand{
		Title: reqBody.Title,
	}
	createdArticle, err := ah.articleUsecase.CreateArticle(command)
	if err != nil {
		logger.Logger.Error("Handler: usecase.CreateArticle failed", "error", err)
		ctx.Error(err)
		return
	}

	logger.Logger.Info("Handler: CreateArticle succeeded", "id", createdArticle.ID)

	ctx.JSON(http.StatusCreated, gin.H{
		"id":    createdArticle.ID,
		"title": createdArticle.Title,
	})
}

func (ah *ArticleHandler) GetArticleById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Logger.Warn("Handler: invalid article ID", "id", idStr, "error", err)
		ctx.Error(common.NewAppError(common.ErrInvalid, "invalid article ID"))
		return
	}

	logger.Logger.Info("Handler: GetArticleById called", "id", id)

	command := usecase.GetArticleByIdCommand{
		ID: id,
	}
	article, err := ah.articleUsecase.GetArticleById(command)
	if err != nil {
		logger.Logger.Error("Handler: usecase.GetArticleById failed", "id", id, "error", err)
		ctx.Error(err)
		return
	}

	logger.Logger.Info("Handler: GetArticleById succeeded", "id", article.ID)

	ctx.JSON(http.StatusOK, gin.H{
		"id":    article.ID,
		"title": article.Title,
	})
}
