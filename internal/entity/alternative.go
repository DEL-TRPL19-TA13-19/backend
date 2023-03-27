package entity

import (
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	"ta13-svc/pkg/utils/constant"
	"ta13-svc/pkg/utils/date"
)

type AlternativeEntity struct {
	Nama           string `json:"nama"`
	Lokasi         string `json:"lokasi"`
	JarakTpa       string `json:"jarak_tpa"`
	JarakTps       string `json:"jarak-tps"`
	JumlahPenduduk string `json:"jumlah_penduduk"`
	TimbulanSampah string `json:"timbulan_sampah"`
	Aksesibilitas  string `json:"Aksesibilitas"`
	Estetika       string `json:"estetika"`
}

type AlternativeEntityModel struct {
	abstraction.Entity
	AlternativeEntity
	CollectionID string `json:"collection_id" gorm:"size:191"`
	//Context      *abstraction.Context `json:"-" gorm:"-"`
}

func (AlternativeEntityModel) TableName() string {
	return "alternatives"
}

func (m *AlternativeEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DbDefaultCreateBy
	return
}

func (m *AlternativeEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
