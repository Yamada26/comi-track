package handler

import (
	"comi-track/internal/usecase"
	"comi-track/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventUsecase interface {
	CreateEvent(command usecase.CreateEventCommand) (*usecase.EventDTO, error)
}

type EventHandler struct {
	eventUsecase EventUsecase
}

func NewEventHandler(eu EventUsecase) *EventHandler {
	return &EventHandler{eventUsecase: eu}
}

func (eh *EventHandler) CreateEvent(ctx *gin.Context) {
	var reqBody struct {
		Number int `json:"number"`
	}

	// リクエスト受信ログ
	logger.Logger.Info("Handler: CreateEvent called")

	// リクエストボディのバリデーション
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		logger.Logger.Warn("Handler: invalid request body", "error", err)
		ctx.Error(err)
		return
	}

	// ユースケースを呼び出し
	command := usecase.CreateEventCommand{
		Number: reqBody.Number,
	}
	createdEvent, err := eh.eventUsecase.CreateEvent(command)
	if err != nil {
		ctx.Error(err)
		return
	}

	logger.Logger.Info("Handler: CreateEvent succeeded", "number", createdEvent.Number)

	ctx.JSON(http.StatusCreated, gin.H{
		"id":     createdEvent.ID,
		"number": createdEvent.Number,
	})
}
