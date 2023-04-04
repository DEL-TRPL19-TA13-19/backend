package entity

import (
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	"ta13-svc/pkg/utils/constant"
	"ta13-svc/pkg/utils/date"
)

type AlternativeEntity struct {
	Nama                  string `json:"nama" example:"nama"`
	TimbulanSampah        string `json:"timbulan_sampah" example:"Jaringan Jalan"`
	JarakTpa              string `json:"jarak_tpa" example:"Alternatif berada di jangkauan layanan TPA"`
	KondisiTanah          string `json:"kondisi_tanah" example:"Tanah keras tidak memiliki unsur organik dan unsur hara dan kedap air"`
	JarakPemukiman        string `json:"jarak_pemukiman" example:"0m-100m"`
	JarakSungai           string `json:"jarak_sungai" example:"Lokasi memenuhi peil banjir"`
	PartisipasiMasyarakat string `json:"partisipasi_masyarakat" example:"< 20% Masyarakat Setuju"`
	CakupanRumah          string `json:"cakupan_rumah" example:"<40 Rumah"`
	Aksesibilitas         string `json:"aksesibilitas" example:"Kondisi jalan bagus dan bisa dilewati kendaraan pengangkut sampah"`
	Sort                  int8   `json:"sort"swag`
}

type AlternativeScores struct {
	TimbulanSampah        float64
	JarakTpa              float64
	KondisiTanah          float64
	JarakPemukiman        float64
	JarakSungai           float64
	PartisipasiMasyarakat float64
	CakupanRumah          float64
	Aksesibilitas         float64
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
