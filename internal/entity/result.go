package entity

import (
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	"ta13-svc/pkg/utils/constant"
	"ta13-svc/pkg/utils/date"
)

type ResultEntityModel struct {
	abstraction.Entity
	CollectionID  string `json:"collection_id" gorm:"size:191"`
	AlternativeID string `json:"alternative_id" gorm:"size:191"`
}

func (ResultEntityModel) TableName() string {
	return "results"
}

func (m *ResultEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DbDefaultCreateBy
	return
}

func (m *ResultEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
