package entity

import (
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	"ta13-svc/pkg/utils/constant"
	"ta13-svc/pkg/utils/date"
)

type FinalScoreEntity struct {
	FinalScore float64 `json:"final_score"`
	Rank       int8    `json:"rank"`
}

type FinalScoreEntityModel struct {
	abstraction.Entity
	FinalScoreEntity
	AlternativeID string `json:"alternative_id" gorm:"size:191"`
	CollectionID  string `json:"collection_id" gorm:"size:191"`
}

func (FinalScoreEntityModel) TableName() string {
	return "final_scores"
}

func (m *FinalScoreEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DbDefaultCreateBy
	return
}

func (m *FinalScoreEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	return
}
