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
	FindValueAlternativeByCollectionID(ctx context.Context, collectionID *string) ([]entity.AlternativeEntityModel, error)
	FindScoreByCollectionID(ctx context.Context, collectionID *string) ([]entity.ScoreEntityModel, error)
	FindSFinalScoreByCollectionID(ctx context.Context, collectionID *string) ([]entity.FinalScoreEntityModel, error)
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

func (a *ahp) FindValueAlternativeByCollectionID(ctx context.Context, collectionID *string) ([]entity.AlternativeEntityModel, error) {
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

func (a *ahp) FindSFinalScoreByCollectionID(ctx context.Context, collectionID *string) ([]entity.FinalScoreEntityModel, error) {
	var datas []entity.FinalScoreEntityModel

	err := a.Db.Where("collection_id = ?", collectionID).Find(&datas).WithContext(ctx).Error

	if err != nil {
		return datas, err
	}

	return datas, nil
}
