package repository

import (
	"context"
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	"ta13-svc/internal/entity"
)

type CollectionRepository interface {
	FindByID(ctx context.Context, id *string) (*entity.CollectionEntityModel, error)
	FindByUserID(ctx context.Context, userID *string) (*[]entity.CollectionEntityModel, error)
	Create(ctx context.Context, e *entity.CollectionEntityModel) (*entity.CollectionEntityModel, error)
	Update(ctx context.Context, id *string, e *entity.CollectionEntityModel) (*entity.CollectionEntityModel, error)
	Delete(ctx context.Context, id *string, e *entity.CollectionEntityModel) (*entity.CollectionEntityModel, error)
}

type collection struct {
	abstraction.Repository
}

func NewCollection(db *gorm.DB) *collection {
	return &collection{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (c *collection) FindByID(ctx context.Context, id *string) (*entity.CollectionEntityModel, error) {

	var data entity.CollectionEntityModel

	err := c.Db.Where("id = ?", id).First(&data).
		WithContext(ctx).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *collection) FindByUserID(ctx context.Context, userID *string) (*[]entity.CollectionEntityModel, error) {

	var datas []entity.CollectionEntityModel

	//err := conn.Preload("Users").Find(&datas).
	err := c.Db.Where("user_id = ?", userID).Find(&datas).
		WithContext(ctx).Limit(20).Offset(40).Error

	if err != nil {
		return nil, err
	}

	return &datas, nil
}

func (c *collection) Create(ctx context.Context, e *entity.CollectionEntityModel) (*entity.CollectionEntityModel, error) {

	err := c.Db.Create(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	err = c.Db.Model(e).First(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (c *collection) Update(ctx context.Context, id *string, e *entity.CollectionEntityModel) (*entity.CollectionEntityModel, error) {

	err := c.Db.Where("id = ?", id).First(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	err = c.Db.Model(e).Updates(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (c *collection) Delete(ctx context.Context, id *string, e *entity.CollectionEntityModel) (*entity.CollectionEntityModel, error) {
	err := c.Db.Where("id = ?", id).Delete(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}
