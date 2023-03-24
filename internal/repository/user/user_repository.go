package repository

import (
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	"ta13-svc/internal/entity"
)

type UserRepository interface {
	GetAll(ctx *abstraction.Context) (*[]entity.UserEntityModel, error)
	FindByUsername(ctx *abstraction.Context, username *string) (*entity.UserEntityModel, error)
	Create(ctx *abstraction.Context, data *entity.UserEntity) (*entity.UserEntityModel, error)
	checkTrx(ctx *abstraction.Context) *gorm.DB
}

type user struct {
	abstraction.Repository
}

func (u user) GetAll(ctx *abstraction.Context) (*[]entity.UserEntityModel, error) {
	//TODO implement me
	panic("implement me")
}

func (u user) FindByUsername(ctx *abstraction.Context, username *string) (*entity.UserEntityModel, error) {
	//TODO implement me
	panic("implement me")
}

func (u user) Create(ctx *abstraction.Context, data *entity.UserEntity) (*entity.UserEntityModel, error) {
	//TODO implement me
	panic("implement me")
}

func (u user) checkTrx(ctx *abstraction.Context) *gorm.DB {
	//TODO implement me
	panic("implement me")
}

func NewUser(db *gorm.DB) *user {
	return &user{
		abstraction.Repository{
			Db: db,
		},
	}
}
