package domain

import "comi-track/internal/common"

type Event struct {
	id     int
	number int
}

func NewEvent(id int, number int) (*Event, error) {
	if id < 0 {
		return nil, common.NewAppError(common.ErrInvalid, "event id must be non-negative")
	}

	if number <= 0 {
		return nil, common.NewAppError(common.ErrInvalid, "event number must be positive")
	}
	return &Event{id: id, number: number}, nil
}

func (event *Event) GetID() int {
	return event.id
}

func (event *Event) GetNumber() int {
	return event.number
}

type EventRepository interface {
	Create(event *Event) (*Event, error)
}
