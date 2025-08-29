package handler

import (
	"comi-track/internal/domain"
	"comi-track/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleUsecase interface {
	GetArticleById(id int) (*domain.Article, error)
	CreateArticle(article *domain.Article) (*domain.Article, error)
}

type ArticleHandler struct {
	articleUsecase ArticleUsecase
}

func NewArticleHandler(au ArticleUsecase) *ArticleHandler {
	return &ArticleHandler{articleUsecase: au}
}

func (ah *ArticleHandler) CreateArticle(ctx *gin.Context) {
	var reqBody struct {
		Title string `json:"title" binding:"required"`
	}

	// リクエスト受信ログ
	logger.Logger.Info("Handler: CreateArticle called")

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		logger.Logger.Warn("Handler: invalid request body", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	articleToCreate, err := domain.NewArticle(0, reqBody.Title)
	if err != nil {
		logger.Logger.Warn("Handler: failed to create domain.Article", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdArticle, err := ah.articleUsecase.CreateArticle(articleToCreate)
	if err != nil {
		logger.Logger.Error("Handler: usecase.CreateArticle failed", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Logger.Info("Handler: CreateArticle succeeded", "id", createdArticle.GetId())

	ctx.JSON(http.StatusCreated, gin.H{
		"id":    createdArticle.GetId(),
		"title": createdArticle.GetTitle(),
	})
}

func (ah *ArticleHandler) GetArticleById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Logger.Warn("Handler: invalid article ID", "id", idStr, "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	logger.Logger.Info("Handler: GetArticleById called", "id", id)

	article, err := ah.articleUsecase.GetArticleById(id)
	if err != nil {
		logger.Logger.Error("Handler: usecase.GetArticleById failed", "id", id, "error", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	logger.Logger.Info("Handler: GetArticleById succeeded", "id", article.GetId())

	ctx.JSON(http.StatusOK, gin.H{
		"id":    article.GetId(),
		"title": article.GetTitle(),
	})
}
