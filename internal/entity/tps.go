package entity

import (
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	"ta13-svc/pkg/utils/constant"
	"ta13-svc/pkg/utils/date"
)

type TpsEntity struct {
	Nama       string `json:"name"`
	Lokasi     string `json:"location"`
	Kelurahan  string `json:"kelurahan"`
	Kecamatan  string `json:"kecamatan"`
	Kabupaten  string `json:"kabupaten"`
	JarakTPA   string `json:"jarak_tpa"`
	Lattitude  string `json:"lattitude"`
	Longtitude string `json:"longtitude"`
}

type TpsFilter struct {
	Nama   *string `query:"Nama" filter:"ILIKE"`
	Lokasi *string `query:"lokasi"`
}

type TpsEntityModel struct {
	abstraction.Entity
	TpsEntity
	//Context *abstraction.Context `json:"-" gorm:"-"`
}

type TpsFilterModel struct {
	abstraction.Filter
	TpsFilter
}

func (TpsEntityModel) TableName() string {
	return "tps"
}

func (m TpsEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DbDefaultCreateBy
	return
}

func (m *TpsEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
