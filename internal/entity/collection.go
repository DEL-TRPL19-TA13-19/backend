package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	"ta13-svc/pkg/utils/constant"
	"ta13-svc/pkg/utils/date"
)

type CollectionEntity struct {
	Nama      string `json:"nama" example:"Mencari lokasi TPS di balige"`
	Kelurahan string `json:"kelurahan"`
	Kecamatan string `json:"kecamatan"`
	Kabupaten string `json:"kabupaten"`
}

type CollectionEntityModel struct {
	abstraction.Entity
	CollectionEntity
	Alternatives []AlternativeEntityModel `json:"alternatives" gorm:"foreignKey:CollectionID"`
	UserID       uuid.UUID                `json:"user_id" gorm:"size:191"`
	//Context      *abstraction.Context     `json:"-" gorm:"-"`
}

func (CollectionEntityModel) TableName() string {
	return "collections"
}

func (m *CollectionEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DbDefaultCreateBy
	return
}

func (m *CollectionEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
