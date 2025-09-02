package main

import (
	"comi-track/internal/infrastructure/sqlite"
	"comi-track/internal/presentation/gin/handler"
	"comi-track/internal/presentation/gin/middleware"
	"comi-track/internal/usecase"
	"comi-track/pkg/logger"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(middleware.SlogMiddleware(logger.Logger))
	router.Use(middleware.ErrorHandler())
	router.Use(gin.Recovery()) // panic は標準で Recovery してくれる

	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "successful"})
	})

	db, err := sqlite.InitDB("./app.db")
	if err != nil {
		logger.Logger.Error("Failed to initialize database", "error", err)
		log.Fatal(err)
	}

	// article
	articleRepository := sqlite.NewArticleRepository(db)
	articleUsecase := usecase.NewArticleUsecase(articleRepository)
	articleHandler := handler.NewArticleHandler(articleUsecase)

	router.POST("/articles", articleHandler.CreateArticle)
	router.GET("/articles/:id", articleHandler.GetArticleById)

	// event
	eventRepository := sqlite.NewEventRepository(db)
	eventUsecase := usecase.NewEventUsecase(eventRepository)
	eventHandler := handler.NewEventHandler(eventUsecase)

	router.POST("/events", eventHandler.CreateEvent)
	// router.GET("/events/:id", eventHandler.GetArticleById)

	if err := router.Run(":8080"); err != nil {
		logger.Logger.Error("Failed to start server", "error", err)
		log.Fatal(err)
	}
}
