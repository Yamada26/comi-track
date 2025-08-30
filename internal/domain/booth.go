package domain

import (
	"errors"
	"fmt"
	"regexp"
)

// HallName
type HallName struct {
	value string
}

func NewHallName(value string) (HallName, error) {
	if value != "N" && value != "S" {
		return HallName{}, errors.New("hall name must be N or S")
	}

	return HallName{value}, nil
}

func (hallName HallName) GetValue() string {
	return hallName.value
}

// BlockCode
type BlockCode struct {
	value string
}

func NewBlockCode(value string) (BlockCode, error) {
	// 正規表現: 半角英字、ひらがな、カタカナのみ
	var validCode = regexp.MustCompile(`^[A-Za-zぁ-んァ-ン]+$`)

	if !validCode.MatchString(value) {
		return BlockCode{}, fmt.Errorf("invalid block code: %s", value)
	}

	return BlockCode{value}, nil
}

func (blockCode BlockCode) GetValue() string {
	return blockCode.value
}

// SpaceHalf
type SpaceHalf struct {
	value string
}

func NewSpaceHalf(value string) (SpaceHalf, error) {
	if value != "a" && value != "b" {
		return SpaceHalf{}, errors.New("space half must be 'a' or 'b'")
	}

	return SpaceHalf{value}, nil
}

// BoothLocation
type BoothLocation struct {
	hallName    HallName
	blockCode   BlockCode
	spaceNumber int
	spaceHalf   SpaceHalf
}

func NewBoothLocation(hallName HallName, blockCode BlockCode, spaceNumber int, spaceHalf SpaceHalf) (BoothLocation, error) {
	if spaceNumber <= 0 || spaceNumber >= 100 {
		return BoothLocation{}, errors.New("space number must be in the range 1 to 99")
	}

	return BoothLocation{hallName, blockCode, spaceNumber, spaceHalf}, nil
}

// BoothId
type BoothId struct {
	value int
}

func NewBoothId(value int) (BoothId, error) {
	if value < 0 {
		return BoothId{}, errors.New("booth id must be positive or 0")
	}

	return BoothId{value}, nil
}

// Booth
type Booth struct {
	id          BoothId
	eventNumber EventNumber
	day         int
	location    BoothLocation
}

func NewBooth(id int, eventNumber EventNumber, day int, location BoothLocation) (*Booth, error) {
	if id < 0 {
		return &Booth{}, errors.New("booth id must be positive or 0")
	}

	boothId, err := NewBoothId(id)
	if err != nil {
		return &Booth{}, err
	}

	return &Booth{
		id:          boothId,
		eventNumber: eventNumber,
		day:         day,
		location:    location,
	}, nil
}

type BoothRepository interface {
	Create(booth *Booth) (*Booth, error)
	FindAll() ([]*Booth, error)
}
