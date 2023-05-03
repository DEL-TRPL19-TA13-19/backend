package collection

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	dto "ta13-svc/internal/dto/collection"
	"ta13-svc/internal/entity"
	"ta13-svc/internal/factory"
	"ta13-svc/internal/repository"
	"ta13-svc/pkg/response"
	"ta13-svc/pkg/utils/trxmanager"
)

type Service interface {
	FindAll(ctx context.Context) ([]entity.CollectionEntityModel, error)
	FindByID(ctx context.Context, payload *dto.CollectionGetByIDRequest) (*dto.CollectionGetByIDResponse, error)
	FindByUserId(ctx context.Context, payload *dto.CollectionsGetByUserIDRequest) ([]entity.CollectionEntityModel, error)
	Create(ctx context.Context, payload *dto.CollectionCreateRequest) (*dto.CollectionCreateResponse, error)
	Update(ctx context.Context, payload *dto.CollectionUpdateRequest) (*dto.CollectionUpdateResponse, error)
	Delete(ctx context.Context, payload *dto.CollectionDeleteRequest) (*dto.CollectionDeleteResponse, error)
}

type service struct {
	Repository repository.CollectionRepository
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.CollectionRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) FindAll(ctx context.Context) ([]entity.CollectionEntityModel, error) {
	datas := make([]entity.CollectionEntityModel, 0)

	datas, err = s.Repository.FindAll(ctx)
	if err != nil {
		return datas, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	return datas, nil
}

func (s *service) FindByUserId(ctx context.Context, payload *dto.CollectionsGetByUserIDRequest) ([]entity.CollectionEntityModel, error) {
	datas := make([]entity.CollectionEntityModel, 0)

	datas, err = s.Repository.FindByUserID(ctx, &payload.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return datas, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}
		return datas, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	return datas, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.CollectionGetByIDRequest) (*dto.CollectionGetByIDResponse, error) {
	var result *dto.CollectionGetByIDResponse

	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}
		return result, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	result = &dto.CollectionGetByIDResponse{
		Datas: *data,
	}

	return result, nil
}

func (s *service) Create(ctx context.Context, payload *dto.CollectionCreateRequest) (*dto.CollectionCreateResponse, error) {
	var result *dto.CollectionCreateResponse
	var data *entity.CollectionEntityModel

	var userID string = ctx.Value("user").(string)
	uuID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	if err = trxmanager.New(s.Db).WithTrxV2(ctx, func(ctx context.Context, f *factory.Factory) error {
		collectionRepository := f.CollectionRepository
		data = &entity.CollectionEntityModel{CollectionEntity: payload.CollectionEntity, UserID: uuID, Entity: abstraction.Entity{
			ID: uuid.NewString(),
		}}
		_, err := collectionRepository.Create(ctx, data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.CollectionCreateResponse{
		CollectionEntityModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx context.Context, payload *dto.CollectionUpdateRequest) (*dto.CollectionUpdateResponse, error) {
	var result *dto.CollectionUpdateResponse
	var data *entity.CollectionEntityModel

	if err = trxmanager.New(s.Db).WithTrxV2(ctx, func(ctx context.Context, f *factory.Factory) error {
		collectionRepository := f.CollectionRepository

		data = &entity.CollectionEntityModel{
			Entity:           abstraction.Entity{ID: payload.ID},
			CollectionEntity: payload.CollectionEntity,
		}
		_, err := collectionRepository.FindByID(ctx, &payload.ID)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}

		_, err = collectionRepository.Update(ctx, &payload.ID, data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
		}
		return nil

	}); err != nil {
		return result, err
	}

	result = &dto.CollectionUpdateResponse{
		CollectionEntityModel: *data,
	}

	return result, nil
}

func (s *service) Delete(ctx context.Context, payload *dto.CollectionDeleteRequest) (*dto.CollectionDeleteResponse, error) {
	var result *dto.CollectionDeleteResponse
	var data *entity.CollectionEntityModel

	if err = trxmanager.New(s.Db).WithTrxV2(ctx, func(ctx context.Context, f *factory.Factory) error {
		collectionRepository := f.CollectionRepository

		data = &entity.CollectionEntityModel{
			CollectionEntity: payload.CollectionEntity,
		}

		_, err = collectionRepository.FindByID(ctx, &payload.ID)

		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}

		_, err = collectionRepository.Delete(ctx, &payload.ID, data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.CollectionDeleteResponse{
		ID: &payload.ID,
	}

	return result, nil
}
