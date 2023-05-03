package entity

import (
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	"ta13-svc/pkg/utils/constant"
	"ta13-svc/pkg/utils/date"
)

type ScoreEntity struct {
	TimbulanSampah        float64 `json:"timbulan_sampah"`
	JarakTpa              float64 `json:"jarak_tpa"`
	JarakPemukiman        float64 `json:"jarak_pemukiman"`
	JarakSungai           float64 `json:"jarak_sungai"`
	PartisipasiMasyarakat float64 `json:"partisipasi_masyarakat"`
	CakupanRumah          float64 `json:"cakupan_rumah"`
	Aksesibilitas         float64 `json:"aksesibilitas"`
}

type ScoreEntityModel struct {
	abstraction.Entity
	ScoreEntity
	AlternativeID string `json:"alternative_id" gorm:"size:191"`
	CollectionID  string `json:"collection_id" gorm:"size:191"`
}

func (ScoreEntityModel) TableName() string {
	return "scores"
}

func (m *ScoreEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DbDefaultCreateBy
	return
}

func (m *ScoreEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	return
}
