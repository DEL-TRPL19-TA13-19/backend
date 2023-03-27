package repository

import (
	"context"
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	"ta13-svc/internal/entity"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username *string) (*entity.UserEntityModel, error)
	Create(ctx context.Context, e *entity.UserEntity) (*entity.UserEntityModel, error)
}

type user struct {
	abstraction.Repository
}

func NewUser(db *gorm.DB) *user {
	return &user{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (u user) FindByUsername(ctx context.Context, username *string) (*entity.UserEntityModel, error) {
	var data entity.UserEntityModel

	err := u.Db.Where("username = ?", username).First(&data).WithContext(ctx).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (u user) Create(ctx context.Context, e *entity.UserEntity) (*entity.UserEntityModel, error) {
	var data entity.UserEntityModel
	data.Entity.BeforeCreate(u.Db)
	data.UserEntity = *e

	err := u.Db.Create(&data).WithContext(ctx).Error

	if err != nil {
		return nil, err
	}

	err = u.Db.Model(&data).First(&data).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}
