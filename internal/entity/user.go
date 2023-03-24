package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	"ta13-svc/pkg/constant"
	"ta13-svc/pkg/date"
)

type UserEntity struct {
	FullName     string
	Username     string `json:"username"`
	Email        string
	Password     string `gorm:"-"`
	PasswordHash string `json:"-"`
	isActive     bool
}
type UserEntityModel struct {
	abstraction.Entity
	UserEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (m UserEntityModel) TableName() string {
	return "users"
}

func (m UserEntityModel) hashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	m.PasswordHash = string(bytes)
}

//func (m UserEntityModel) GenerateToken() (string, error) {
////	TODO : membuat generate token JWT
//	return
//}

func (m *UserEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DbDefaultCreateBy
	m.hashPassword()
	m.Password = ""
	return
}

func (m *UserEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
