package repository

import (
	"context"
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	"ta13-svc/internal/entity"
)

type TpsRepository interface {
	FindAll(ctx context.Context) (*[]entity.TpsEntityModel, error)
	FindByID(ctx context.Context, id *string) (*entity.TpsEntityModel, error)
	Create(ctx context.Context, m *entity.TpsEntityModel) (*entity.TpsEntityModel, error)
	Update(ctx context.Context, id *string, m *entity.TpsEntityModel) (*entity.TpsEntityModel, error)
	Delete(ctx context.Context, id *string, m *entity.TpsEntityModel) (*entity.TpsEntityModel, error)
}

type tps struct {
	abstraction.Repository
}

func NewTps(db *gorm.DB) *tps {
	return &tps{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (t *tps) FindAll(ctx context.Context) (*[]entity.TpsEntityModel, error) {

	var datas []entity.TpsEntityModel

	err := t.Db.Find(&datas).
		WithContext(ctx).Error
	if err != nil {
		return &datas, err
	}

	return &datas, nil
}

func (t *tps) FindByID(ctx context.Context, id *string) (*entity.TpsEntityModel, error) {

	var data entity.TpsEntityModel
	err := t.Db.Where("id = ?", id).First(&data).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (t *tps) Create(ctx context.Context, e *entity.TpsEntityModel) (*entity.TpsEntityModel, error) {

	err := t.Db.Create(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	err = t.Db.Model(e).First(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (t *tps) Update(ctx context.Context, id *string, e *entity.TpsEntityModel) (*entity.TpsEntityModel, error) {
	err := t.Db.Model(e).Where("id = ?", id).Updates(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	print(e)
	return e, nil
}

func (t *tps) Delete(ctx context.Context, id *string, e *entity.TpsEntityModel) (*entity.TpsEntityModel, error) {
	err := t.Db.Where("id = ?", id).Delete(e).
		WithContext(ctx).Error

	if err != nil {
		return nil, err
	}

	return e, nil
}
