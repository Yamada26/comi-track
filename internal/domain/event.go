package domain

import "comi-track/internal/common"

type EventNumber struct {
	value int
}

func NewEventNumber(value int) (EventNumber, error) {
	if value <= 0 {
		return EventNumber{}, common.NewAppError(common.ErrInvalid, "event number must be positive")
	}

	return EventNumber{value}, nil
}

type Event struct {
	number EventNumber
}

func NewEvent(number EventNumber) (*Event, error) {
	return &Event{number: number}, nil
}

func (event *Event) GetNumber() EventNumber {
	return event.number
}
