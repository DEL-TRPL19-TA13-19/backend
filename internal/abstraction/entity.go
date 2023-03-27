package abstraction

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ta13-svc/pkg/utils/date"
	"time"
)

type Entity struct {
	ID string `json:"id" gorm:"primaryKey"`

	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  string     `json:"created_by"`
	ModifiedAt *time.Time `json:"modified_at"`
	ModifiedBy *string    `json:"modified_by"`
}

type Filter struct {
	CreatedAt  *time.Time `query:"created_at"`
	CreatedBy  *string    `query:"created_by"`
	ModifiedAt *time.Time `query:"modified_at"`
	ModifiedBy *string    `query:"modified_by"`
}

func (e *Entity) BeforeCreate(trx *gorm.DB) (err error) {
	e.CreatedAt = *date.DateTodayLocal()
	e.ID = uuid.NewString()
	return
}

func (e *Entity) BeforeUpdate(trx *gorm.DB) (err error) {
	e.ModifiedAt = date.DateTodayLocal()
	return
}
