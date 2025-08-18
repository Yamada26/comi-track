package presentation

import (
	"comi-track/infrastructure"
	"comi-track/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	articleRepository := infrastructure.NewArticleRepository()
	articleUsecase := usecase.NewArticleUsecase(articleRepository)
	articleHandler := NewArticleHandler(articleUsecase)

	router.POST("/articles", articleHandler.CreateArticle)
	router.GET("/articles/:id", articleHandler.GetArticleById)

	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "successful"})
	})

	return router
}
