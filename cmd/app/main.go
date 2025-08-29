package main

import (
	infrastructure "comi-track/internal/infrastructure/sqlite"
	presentation "comi-track/internal/presentation/gin"
	"comi-track/internal/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := infrastructure.InitDB("./app.db")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "successful"})
	})

	articleRepository := infrastructure.NewArticleRepository(db)
	articleUsecase := usecase.NewArticleUsecase(articleRepository)
	articleHandler := presentation.NewArticleHandler(articleUsecase)

	router.POST("/articles", articleHandler.CreateArticle)
	router.GET("/articles/:id", articleHandler.GetArticleById)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
