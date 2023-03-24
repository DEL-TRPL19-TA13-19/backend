package auth

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"ta13-svc/internal/abstraction"
	dto "ta13-svc/internal/dto/auth"
	"ta13-svc/internal/entity"
	"ta13-svc/internal/factory"
	repository "ta13-svc/internal/repository/user"
	"ta13-svc/pkg/response"
	"ta13-svc/pkg/trxmanager"
)

type Service interface {
	Login(ctx *abstraction.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	Register(ctx *abstraction.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error)
}

type service struct {
	Repository repository.UserRepository
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.UserRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Login(c *abstraction.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	var result *dto.AuthLoginResponse

	data, err := s.Repository.FindByUsername(c, &payload.Username)

	if data == nil {
		return result, response.ErrorBuilder(&response.ErrorConstant.Unauthorized, err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data.PasswordHash), []byte(payload.Password)); err != nil {
		return result, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	result = &dto.AuthLoginResponse{
		Message:         "Berhasil login",
		UserEntityModel: *data,
	}

	return result, nil
}

func (s *service) Register(c *abstraction.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error) {
	var result *dto.AuthRegisterResponse
	var data *entity.UserEntityModel

	if err = trxmanager.New(s.Db).WithTrx(c, func(c *abstraction.Context) error {
		data, err = s.Repository.Create(c, &payload.UserEntity)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.AuthRegisterResponse{
		UserEntityModel: *data,
	}

	return result, nil
}
