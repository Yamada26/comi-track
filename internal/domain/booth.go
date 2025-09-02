package domain

import (
	"comi-track/internal/common"
	"fmt"
	"regexp"
)

// BoothLocation
type BoothLocation struct {
	hallName    string
	blockCode   string
	spaceNumber int
	spaceHalf   string
}

func NewBoothLocation(hallName string, blockCode string, spaceNumber int, spaceHalf string) (BoothLocation, *common.AppError) {
	if hallName != "N" && hallName != "S" {
		return BoothLocation{}, common.NewAppError(common.ErrInvalid, "hall name must be 'N' or 'S'")
	}

	// 正規表現: 半角英字、ひらがな、カタカナのみ
	var validCode = regexp.MustCompile(`^[A-Za-zぁ-んァ-ン]+$`)
	if !validCode.MatchString(blockCode) {
		return BoothLocation{}, common.NewAppError(common.ErrInvalid, fmt.Sprintf("block code '%s' is invalid", blockCode))
	}

	if spaceNumber <= 0 || spaceNumber >= 100 {
		return BoothLocation{}, common.NewAppError(common.ErrInvalid, "space number must be between 1 and 99")
	}

	if spaceHalf != "a" && spaceHalf != "b" {
		return BoothLocation{}, common.NewAppError(common.ErrInvalid, "space half must be 'a' or 'b'")
	}

	return BoothLocation{hallName, blockCode, spaceNumber, spaceHalf}, nil
}

func (location *BoothLocation) GetHallName() string {
	return location.hallName
}

func (location *BoothLocation) GetBlockCode() string {
	return location.blockCode
}

func (location *BoothLocation) GetSpaceNumber() int {
	return location.spaceNumber
}

func (location *BoothLocation) GetSpaceHalf() string {
	return location.spaceHalf
}

// BoothId
type BoothId struct {
	value int
}

func NewBoothId(value int) (BoothId, *common.AppError) {
	if value < 0 {
		return BoothId{}, common.NewAppError(common.ErrInvalid, "booth id must be non-negative")
	}

	return BoothId{value}, nil
}

func (boothId BoothId) GetValue() int {
	return boothId.value
}

// Booth
type Booth struct {
	id          BoothId
	eventNumber int
	day         int
	location    BoothLocation
}

func NewBooth(id BoothId, eventNumber int, day int, location BoothLocation) (*Booth, *common.AppError) {
	return &Booth{
		id:          id,
		eventNumber: eventNumber,
		day:         day,
		location:    location,
	}, nil
}

func (booth *Booth) GetID() BoothId {
	return booth.id
}

func (booth *Booth) GetEventNumber() int {
	return booth.eventNumber
}

func (booth *Booth) GetDay() int {
	return booth.day
}

func (booth *Booth) GetLocation() BoothLocation {
	return booth.location
}

// BoothRepository
type BoothRepository interface {
	Create(booth *Booth) (*Booth, *common.AppError)
	FindAll() ([]*Booth, *common.AppError)
}
