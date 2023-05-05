package repository

import (
	"context"
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	"ta13-svc/internal/entity"
)

type AhpRepository interface {
	CreateScore(ctx context.Context, e []entity.ScoreEntityModel) ([]entity.ScoreEntityModel, error)
	CreateFinalScore(ctx context.Context, e []entity.FinalScoreEntityModel) ([]entity.FinalScoreEntityModel, error)

	FindAlternativesByCollectionID(ctx context.Context, collectionID *string) ([]entity.AlternativeEntityModel, error)
	FindScoreByCollectionID(ctx context.Context, collectionID *string) ([]entity.ScoreEntityModel, error)
	FindFinalScoreByCollectionID(ctx context.Context, collectionID *string) ([]entity.AlternativeEntityModel, error)

	UpdateCollection(ctx context.Context, collectionID *string, e *entity.CollectionEntityModel) (*entity.CollectionEntityModel, error)

	DeleteAllScoreByCollection(ctx context.Context, collectionID *string) (*entity.ScoreEntityModel, error)
	DeleteAllFinalScoreByCollection(ctx context.Context, collectionID *string) (*entity.FinalScoreEntityModel, error)
}

type ahp struct {
	abstraction.Repository
}

func NewAHP(db *gorm.DB) *ahp {
	return &ahp{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (a *ahp) CreateScore(ctx context.Context, e []entity.ScoreEntityModel) ([]entity.ScoreEntityModel, error) {
	err := a.Db.Create(&e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (a *ahp) CreateFinalScore(ctx context.Context, e []entity.FinalScoreEntityModel) ([]entity.FinalScoreEntityModel, error) {
	err := a.Db.Create(&e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (a *ahp) FindAlternativesByCollectionID(ctx context.Context, collectionID *string) ([]entity.AlternativeEntityModel, error) {
	var datas []entity.AlternativeEntityModel

	err := a.Db.Where("collection_id = ?", collectionID).Find(&datas).WithContext(ctx).Error

	if err != nil {
		return datas, err
	}

	return datas, nil
}

func (a *ahp) FindScoreByCollectionID(ctx context.Context, collectionID *string) ([]entity.ScoreEntityModel, error) {
	var datas []entity.ScoreEntityModel

	err := a.Db.Where("collection_id = ?", collectionID).Find(&datas).WithContext(ctx).Error

	if err != nil {
		return datas, err
	}

	return datas, nil
}

func (a *ahp) FindFinalScoreByCollectionID(ctx context.Context, collectionID *string) ([]entity.AlternativeEntityModel, error) {
	var datas []entity.AlternativeEntityModel

	err := a.Db.Preload("FinalScore").Where("collection_id = ?", collectionID).Find(&datas).WithContext(ctx).Error

	if err != nil {
		return datas, err
	}

	return datas, nil
}

func (a *ahp) UpdateCollection(ctx context.Context, collectionID *string, e *entity.CollectionEntityModel) (*entity.CollectionEntityModel, error) {

	err := a.Db.Model(e).Where("id", collectionID).Updates(e).WithContext(ctx).Error

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (a *ahp) DeleteAllScoreByCollection(ctx context.Context, collectionID *string) (*entity.ScoreEntityModel, error) {
	var data *entity.ScoreEntityModel

	err := a.Db.Where("collection_id = ?", collectionID).Delete(&data).WithContext(ctx).Error

	if err != nil {
		return data, err
	}

	return data, err
}

func (a *ahp) DeleteAllFinalScoreByCollection(ctx context.Context, collectionID *string) (*entity.FinalScoreEntityModel, error) {
	var data *entity.FinalScoreEntityModel

	err := a.Db.Where("collection_id = ?", collectionID).Delete(&data).WithContext(ctx).Error

	if err != nil {
		return data, err
	}

	return data, err
}
