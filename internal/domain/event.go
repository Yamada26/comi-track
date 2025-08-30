package domain

import "errors"

type EventNumber struct {
	value int
}

func NewEventNumber(value int) (EventNumber, error) {
	if value <= 0 {
		return EventNumber{}, errors.New("event number bust be positive")
	}

	return EventNumber{value}, nil
}

type Event struct {
	number EventNumber
}

func NewEvent(number EventNumber) (*Event, error) {
	return &Event{number: number}, nil
}
