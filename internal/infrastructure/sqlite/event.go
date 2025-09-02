package sqlite

import (
	"comi-track/internal/common"
	"comi-track/internal/domain"
	"comi-track/pkg/logger"

	"gorm.io/gorm"
)

type EventModel struct {
	ID     int `gorm:"primaryKey;column:id;autoIncrement"`
	Number int `gorm:"column:number"`
}

func (EventModel) TableName() string {
	return "events"
}

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

/*
指定された Event を保存する。
保存に成功した場合は、保存された Event を返す。
*/
func (er *EventRepository) Create(event *domain.Event) (*domain.Event, error) {
	logger.Logger.Info("Repository: Create called", "title", event.GetNumber())

	var createdEvent *domain.Event

	err := er.db.Transaction(func(tx *gorm.DB) error {
		model := EventModel{
			Number: event.GetNumber(),
		}

		if err := tx.Create(&model).Error; err != nil {
			logger.Logger.Error("Repository: failed to insert event", "error", err)
			return common.NewAppError(common.ErrInternal, "failed to create event")
		}

		logger.Logger.Info("Repository: event inserted successfully", "id", model.ID)

		var err error
		createdEvent, err = domain.NewEvent(model.ID, model.Number)
		return err
	})

	if err != nil {
		return nil, err
	}

	return createdEvent, nil
}
