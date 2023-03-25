package repository

import (
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	"ta13-svc/internal/entity"
)

type UserRepository interface {
	FindByUsername(ctx *abstraction.Context, username *string) (*entity.UserEntityModel, error)
	Create(ctx *abstraction.Context, payload *entity.UserEntity) (*entity.UserEntityModel, error)
	checkTrx(ctx *abstraction.Context) *gorm.DB
}

type user struct {
	abstraction.Repository
}

func (u user) FindByUsername(ctx *abstraction.Context, username *string) (*entity.UserEntityModel, error) {
	conn := u.checkTrx(ctx)
	var data entity.UserEntityModel

	err := conn.Where("username = ?", username).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (u user) Create(ctx *abstraction.Context, e *entity.UserEntity) (*entity.UserEntityModel, error) {
	conn := u.checkTrx(ctx)
	var data entity.UserEntityModel
	data.Entity.BeforeCreate(conn)
	data.UserEntity = *e

	err := conn.Create(&data).Error

	if err != nil {
		return nil, err
	}

	err = conn.Model(&data).First(&data).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (u user) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return u.Db
}

func NewUser(db *gorm.DB) *user {
	return &user{
		abstraction.Repository{
			Db: db,
		},
	}
}
