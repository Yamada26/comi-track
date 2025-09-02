package usecase

import (
	"comi-track/internal/domain"
	"comi-track/pkg/logger"
)

type EventUsecase struct {
	eventRepository domain.EventRepository
}

func NewEventUsecase(ar domain.EventRepository) *EventUsecase {
	return &EventUsecase{eventRepository: ar}
}

type EventDTO struct {
	ID     int `json:"id"`
	Number int `json:"number"`
}

type CreateEventCommand struct {
	Number int `json:"number"`
}

func (eu *EventUsecase) CreateEvent(command CreateEventCommand) (*EventDTO, error) {
	logger.Logger.Info("Usecase: CreateEvent called", "number", command.Number)

	eventToCreate, err := domain.NewEvent(0, command.Number)
	if err != nil {
		return nil, err
	}

	createdEvent, err := eu.eventRepository.Create(eventToCreate)
	if err != nil {
		return nil, err
	}

	logger.Logger.Info("Usecase: CreateEvent succeeded", "id", createdEvent.GetNumber())
	return &EventDTO{
		ID:     createdEvent.GetID(),
		Number: createdEvent.GetNumber(),
	}, nil
}
