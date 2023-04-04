package collection

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"os"
	"ta13-svc/internal/abstraction"
	dto "ta13-svc/internal/dto/collection"
	"ta13-svc/internal/entity"
	"ta13-svc/internal/factory"
	"ta13-svc/internal/repository"
	"ta13-svc/pkg/response"
	"ta13-svc/pkg/utils/ahp"
	"ta13-svc/pkg/utils/trxmanager"
)

type Service interface {
	FindAll(ctx context.Context) ([]entity.CollectionEntityModel, error)
	FindByID(ctx context.Context, payload *dto.CollectionGetByIDRequest) (*dto.CollectionGetByIDResponse, error)
	FindByUserId(ctx context.Context, payload *dto.CollectionsGetByUserIDRequest) ([]entity.CollectionEntityModel, error)
	Create(ctx context.Context, payload *dto.CollectionCreateRequest) (*dto.CollectionCreateResponse, error)
	Update(ctx context.Context, payload *dto.CollectionUpdateRequest) (*dto.CollectionUpdateResponse, error)
	Delete(ctx context.Context, payload *dto.CollectionDeleteRequest) (*dto.CollectionDeleteResponse, error)
	CalculateAHP(ctx context.Context, payload *dto.CollectionGetByIDRequest) ([]entity.AlternativeEntityModel, error)
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

type PairwiseMatrixFromJSON struct {
	Pairwise [][]float64 `json:"pairwise"`
}

type Matrix [][]float64

func (s *service) CalculateAHP(ctx context.Context, payload *dto.CollectionGetByIDRequest) ([]entity.AlternativeEntityModel, error) {
	datas := make([]entity.AlternativeEntityModel, 0)

	datas, err = s.Repository.FindAlternatives(ctx, &payload.ID)

	matrix := make(Matrix, 0)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return datas, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}
		return datas, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	var name []string

	for i := 0; i < len(datas); i++ {
		row := [][]float64{
			{ahp.TimbulanSampahSubCriteria()[datas[i].TimbulanSampah],
				ahp.JarakTPASubCriteria()[datas[i].JarakTpa],
				ahp.KondisiTanahSubCriteria()[datas[i].KondisiTanah],
				ahp.JarakPemukimanSubCriteria()[datas[i].JarakPemukiman],
				ahp.JarakSungaiSubCriteria()[datas[i].JarakSungai],
				ahp.PartisipasiMasyarakatSubCriteria()[datas[i].PartisipasiMasyarakat],
				ahp.CakupanRumahSubCriteria()[datas[i].CakupanRumah],
				ahp.AksesibilitasSubCriteria()[datas[i].Aksesibilitas]}}

		name = append(name, datas[i].Nama)
		matrix = append(matrix, row...)
	}

	jsonFile, err := os.ReadFile("./internal/usecase/collection/pairwise.json")
	if err != nil {
		fmt.Println(err)
	}
	var pairwiseComparison PairwiseMatrixFromJSON
	err = json.Unmarshal(jsonFile, &pairwiseComparison)
	if err != nil {
		fmt.Println(err)
	}

	//MENCARI SUM DARI MASING MASING COL
	rowsPC := len(pairwiseComparison.Pairwise)
	colsPC := len(pairwiseComparison.Pairwise[0])
	colSum := make([]float64, len(pairwiseComparison.Pairwise))

	for i := 0; i < rowsPC; i++ {
		for j := 0; j < colsPC; j++ {
			colSum[i] += pairwiseComparison.Pairwise[j][i]
		}
	}

	//NORMALISASI MATRIKS PAIRWISE
	for i := 0; i < rowsPC; i++ {
		for j := 0; j < colsPC; j++ {
			pairwiseComparison.Pairwise[i][j] /= colSum[j]
		}
	}

	//MENCARI JUMLAH NILAI BARIS DAN KOLOM & MENCARI RATA RATA (BOBOT KRITERIA)
	normalColSum := make([]float64, len(pairwiseComparison.Pairwise))
	normalRowSum := make([]float64, len(pairwiseComparison.Pairwise))
	criteriaWeights := make([]float64, len(pairwiseComparison.Pairwise))

	for i := 0; i < rowsPC; i++ {
		sum := 0.0
		for j := 0; j < colsPC; j++ {
			sum += pairwiseComparison.Pairwise[i][j]
			normalColSum[i] += pairwiseComparison.Pairwise[j][i]
			normalRowSum[i] += pairwiseComparison.Pairwise[i][j]
			criteriaWeights[i] = sum / float64(8)
		}
	}

	rowsACS := len(matrix)
	colsACS := len(matrix[0])

	//PERKALIAN MATRIKS ALTERNATIF DENGAN MATRIKS BOBOT
	for i := 0; i < rowsACS; i++ {
		for j := 0; j < colsACS; j++ {
			matrix[i][j] *= criteriaWeights[j]
		}
	}

	//MENJUMLAHKAN NILAI SETIAP BARIS (SKOR AKHIR)
	alternativeScores := make([]float64, len(datas))
	for i := 0; i < rowsACS; i++ {
		for j := 0; j < colsACS; j++ {
			alternativeScores[i] += matrix[i][j]
		}
	}

	return datas, nil
}
