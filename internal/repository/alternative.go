package repository

import (
	"context"
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	"ta13-svc/internal/entity"
)

type AlternativeRepository interface {
	FindAll(ctx context.Context) ([]entity.AlternativeEntityModel, error)
	FindByID(ctx context.Context, id *string) (*entity.AlternativeEntityModel, error)
	FindByCollectionID(ctx context.Context, collectionID *string) ([]entity.AlternativeEntityModel, error)
	Create(ctx context.Context, e *entity.AlternativeEntityModel) (*entity.AlternativeEntityModel, error)
	Update(ctx context.Context, id *string, e *entity.AlternativeEntityModel) (*entity.AlternativeEntityModel, error)
	Delete(ctx context.Context, id *string, e *entity.AlternativeEntityModel) (*entity.AlternativeEntityModel, error)
}

type alternative struct {
	abstraction.Repository
}

func NewAlternative(db *gorm.DB) *alternative {
	return &alternative{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (a *alternative) FindAll(ctx context.Context) ([]entity.AlternativeEntityModel, error) {
	var datas []entity.AlternativeEntityModel
	err := a.Db.Find(&datas).WithContext(ctx).Error
	if err != nil {
		return datas, err
	}
	return datas, nil
}

func (a *alternative) FindByID(ctx context.Context, id *string) (*entity.AlternativeEntityModel, error) {

	var data entity.AlternativeEntityModel
	err := a.Db.Where("id = ?", id).First(&data).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (a *alternative) FindByCollectionID(ctx context.Context, collectionID *string) ([]entity.AlternativeEntityModel, error) {

	var datas []entity.AlternativeEntityModel
	err := a.Db.Where("collection_id = ?", collectionID).Find(&datas).
		WithContext(ctx).Error
	if err != nil {
		return datas, err
	}

	return datas, nil
}

func (a *alternative) Create(ctx context.Context, e *entity.AlternativeEntityModel) (*entity.AlternativeEntityModel, error) {
	err := a.Db.Create(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	err = a.Db.Model(e).First(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (a *alternative) Update(ctx context.Context, id *string, e *entity.AlternativeEntityModel) (*entity.AlternativeEntityModel, error) {
	err := a.Db.Model(e).Where("id = ?", id).Updates(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	err = a.Db.Model(e).Updates(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (a *alternative) Delete(ctx context.Context, id *string, e *entity.AlternativeEntityModel) (*entity.AlternativeEntityModel, error) {
	err := a.Db.Where("id = ?", id).Delete(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}
