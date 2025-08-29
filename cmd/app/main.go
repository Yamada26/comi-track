package main

import (
	"comi-track/internal/infrastructure/sqlite"
	"comi-track/internal/presentation/gin/handler"
	"comi-track/internal/presentation/gin/middleware"
	"comi-track/internal/usecase"
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := sqlite.InitDB("./app.db")
	if err != nil {
		log.Fatal(err)
	}



	router := gin.New()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	router.Use(middleware.SlogMiddleware(logger))
    router.Use(gin.Recovery()) // panic は標準で Recovery してくれる

	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "successful"})
	})

	articleRepository := sqlite.NewArticleRepository(db)
	articleUsecase := usecase.NewArticleUsecase(articleRepository)
	articleHandler := handler.NewArticleHandler(articleUsecase)

	router.POST("/articles", articleHandler.CreateArticle)
	router.GET("/articles/:id", articleHandler.GetArticleById)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
