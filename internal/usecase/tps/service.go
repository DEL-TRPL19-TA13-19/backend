package tps

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	"ta13-svc/internal/dto/tps"
	"ta13-svc/internal/entity"
	"ta13-svc/internal/factory"
	"ta13-svc/internal/repository"
	"ta13-svc/pkg/response"
	"ta13-svc/pkg/utils/trxmanager"
)

type Service interface {
	FindAll(ctx context.Context) (*dto.TpsGetResponse, error)
	FindById(ctx context.Context, payload *dto.TpsGetByIdRequest) (*dto.TpsGetByIdResponse, error)
	Create(ctx context.Context, payload *dto.TpsCreateRequest) (*dto.TpsCreateResponse, error)
	Update(ctx context.Context, payload *dto.TpsUpdateRequest) (*dto.TpsUpdateResponse, error)
	Delete(ctx context.Context, payload *dto.TpsDeleteRequest) (*dto.TpsDeleteResponse, error)
}

type service struct {
	Repository repository.TpsRepository
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.TpsRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) FindAll(ctx context.Context) (*dto.TpsGetResponse, error) {
	var result *dto.TpsGetResponse
	var datas *[]entity.TpsEntityModel

	datas, err := s.Repository.FindAll(ctx)

	if err != nil {
		return result, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	result = &dto.TpsGetResponse{
		Datas: *datas,
	}

	return result, nil
}

func (s *service) FindById(ctx context.Context, payload *dto.TpsGetByIdRequest) (*dto.TpsGetByIdResponse, error) {
	var result *dto.TpsGetByIdResponse

	data, err := s.Repository.FindByID(ctx, &payload.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}
		return result, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	result = &dto.TpsGetByIdResponse{
		TpsEntityModel: *data,
	}

	return result, nil
}

func (s *service) Create(ctx context.Context, payload *dto.TpsCreateRequest) (*dto.TpsCreateResponse, error) {
	var result *dto.TpsCreateResponse
	var data *entity.TpsEntityModel

	if err = trxmanager.New(s.Db).WithTrxV2(ctx, func(ctx context.Context, f *factory.Factory) error {
		TpsRepository := f.TpsRepository
		data = &entity.TpsEntityModel{
			Entity:    abstraction.Entity{ID: uuid.NewString()},
			TpsEntity: payload.TpsEntity,
		}
		_, err := TpsRepository.Create(ctx, data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.TpsCreateResponse{
		TpsEntityModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx context.Context, payload *dto.TpsUpdateRequest) (*dto.TpsUpdateResponse, error) {
	var result *dto.TpsUpdateResponse
	var data *entity.TpsEntityModel

	if err = trxmanager.New(s.Db).WithTrxV2(ctx, func(ctx context.Context, f *factory.Factory) error {
		tpsRepository := f.TpsRepository

		data = &entity.TpsEntityModel{
			TpsEntity: payload.TpsEntity,
		}
		//_, err = tpsRepository.FindByID(ctx, &payload.ID)
		//if err != nil {
		//	return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err)
		//}

		_, err = tpsRepository.Update(ctx, &payload.ID, data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.TpsUpdateResponse{
		TpsEntityModel: *data,
	}

	return result, nil
}

func (s *service) Delete(ctx context.Context, payload *dto.TpsDeleteRequest) (*dto.TpsDeleteResponse, error) {
	var result *dto.TpsDeleteResponse
	var data *entity.TpsEntityModel

	if err = trxmanager.New(s.Db).WithTrxV2(ctx, func(ctx context.Context, f *factory.Factory) error {
		tpsRepository := f.TpsRepository
		//data, err = tpsRepository.FindByID(ctx, &payload.ID)
		//if err != nil {
		//	return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err)
		//}
		data = &entity.TpsEntityModel{
			TpsEntity: payload.TpsEntity,
		}
		_, err = tpsRepository.Delete(ctx, &payload.ID, data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.TpsDeleteResponse{
		ID: payload.ID,
	}

	return result, nil
}
