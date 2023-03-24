package abstraction

import (
	"gorm.io/gorm"
	"ta13-svc/pkg/date"
	"time"
)

type Entity struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement;"`

	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  string     `json:"created_by"`
	ModifiedAt *time.Time `json:"modified_at"`
	ModifiedBy *string    `json:"modified_by"`
}

func (e *Entity) BeforeCreate(trx *gorm.DB) (err error) {
	e.CreatedAt = *date.DateTodayLocal()
	return
}

func (e *Entity) BeforeUpdate(trx *gorm.DB) (err error) {
	e.ModifiedAt = date.DateTodayLocal()
	return
}
