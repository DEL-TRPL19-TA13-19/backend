package ahp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"os"
	"ta13-svc/internal/entity"
	"ta13-svc/internal/factory"
	"ta13-svc/internal/repository"
	"ta13-svc/pkg/response"
	"ta13-svc/pkg/utils/ahp"
)

type Service interface {
	FindCriteriaAlternative(ctx context.Context) (*CriteriaData, error)
	CalculateScoreAlternativeByCollectionID(ctx context.Context, collectionID *string) ([]entity.ScoreEntityModel, error)
	CalculateFinalScoreByCollectionID(ctx context.Context, collectionID *string) ([]entity.FinalScoreEntityModel, error)
}

type CriteriaData struct {
	Pairwise [][]float64 `json:"pairwise"`
	Criteria []float64   `json:"criteria"`
}

type Matrix [][]float64

type service struct {
	Repository repository.AhpRepository
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.AHPRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) FindCriteriaAlternative(ctx context.Context) (*CriteriaData, error) {
	var result *CriteriaData

	jsonFile, err := os.ReadFile("./internal/usecase/collection/pairwise.json")
	if err != nil {
		fmt.Println(err)
	}
	var criteriaData CriteriaData
	err = json.Unmarshal(jsonFile, &criteriaData)
	if err != nil {
		fmt.Println(err)
	}

	//MENCARI SUM DARI MASING MASING COL
	rowsPC := len(criteriaData.Pairwise)
	colsPC := len(criteriaData.Pairwise[0])
	colSum := make([]float64, len(criteriaData.Pairwise))

	for i := 0; i < rowsPC; i++ {
		for j := 0; j < colsPC; j++ {
			colSum[i] += criteriaData.Pairwise[j][i]
		}
	}

	//NORMALISASI MATRIKS PAIRWISE
	for i := 0; i < rowsPC; i++ {
		for j := 0; j < colsPC; j++ {
			criteriaData.Pairwise[i][j] /= colSum[j]
		}
	}

	//MENCARI JUMLAH NILAI BARIS DAN KOLOM & MENCARI RATA RATA (BOBOT KRITERIA)
	normalColSum := make([]float64, len(criteriaData.Pairwise))
	normalRowSum := make([]float64, len(criteriaData.Pairwise))
	criteriaWeights := make([]float64, len(criteriaData.Pairwise))

	for i := 0; i < rowsPC; i++ {
		sum := 0.0
		for j := 0; j < colsPC; j++ {
			sum += criteriaData.Pairwise[i][j]
			normalColSum[i] += criteriaData.Pairwise[j][i]
			normalRowSum[i] += criteriaData.Pairwise[i][j]
			criteriaWeights[i] = sum / float64(8)
		}
	}

	result = &CriteriaData{
		Pairwise: criteriaData.Pairwise,
		Criteria: criteriaWeights}

	return result, nil
}

func (s *service) CalculateScoreAlternativeByCollectionID(ctx context.Context, collectionID *string) ([]entity.ScoreEntityModel, error) {
	alternatives := make([]entity.AlternativeEntityModel, 0)
	alternatives, err = s.Repository.FindAlternativesByCollectionID(ctx, collectionID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	matrix := make(Matrix, 0)

	for i := 0; i < len(alternatives); i++ {
		row := [][]float64{
			{ahp.TimbulanSampahSubCriteria()[alternatives[i].TimbulanSampah],
				ahp.JarakTPASubCriteria()[alternatives[i].JarakTpa],
				ahp.KondisiTanahSubCriteria()[alternatives[i].KondisiTanah],
				ahp.JarakPemukimanSubCriteria()[alternatives[i].JarakPemukiman],
				ahp.JarakSungaiSubCriteria()[alternatives[i].JarakSungai],
				ahp.PartisipasiMasyarakatSubCriteria()[alternatives[i].PartisipasiMasyarakat],
				ahp.CakupanRumahSubCriteria()[alternatives[i].CakupanRumah],
				ahp.AksesibilitasSubCriteria()[alternatives[i].Aksesibilitas]}}

		matrix = append(matrix, row...)
	}

	criteriaData, _ := s.FindCriteriaAlternative(ctx)
	criteriaWeights := criteriaData.Criteria

	rowsACS := len(matrix)
	colsACS := len(matrix[0])

	//PERKALIAN MATRIKS ALTERNATIF DENGAN MATRIKS BOBOT
	for i := 0; i < rowsACS; i++ {
		for j := 0; j < colsACS; j++ {
			matrix[i][j] *= criteriaWeights[j]
		}
	}

	scores := make([]entity.ScoreEntityModel, 0)

	for i := 0; i < rowsACS; i++ {
		scores = append(scores, entity.ScoreEntityModel{
			ScoreEntity: entity.ScoreEntity{
				TimbulanSampah:        matrix[i][0],
				JarakTpa:              matrix[i][1],
				KondisiTanah:          matrix[i][2],
				JarakPemukiman:        matrix[i][3],
				JarakSungai:           matrix[i][4],
				PartisipasiMasyarakat: matrix[i][5],
				CakupanRumah:          matrix[i][6],
				Aksesibilitas:         matrix[i][7],
			},
			CollectionID:  alternatives[i].CollectionID,
			AlternativeID: alternatives[i].ID,
		})
	}

	_, err := s.Repository.CreateScore(ctx, scores)

	if err != nil {
		return nil, err
	}

	return scores, nil

}

func (s *service) CalculateFinalScoreByCollectionID(ctx context.Context, collectionID *string) ([]entity.FinalScoreEntityModel, error) {
	alternativeScores, err := s.CalculateScoreAlternativeByCollectionID(ctx, collectionID)

	if err != nil {
		return nil, err
	}

	finalscores := make([]entity.FinalScoreEntityModel, 0)

	for i := 0; i < len(alternativeScores); i++ {
		finalscores[i].CollectionID = alternativeScores[i].CollectionID
		finalscores[i].AlternativeID = alternativeScores[i].AlternativeID
		finalscores[i].FinalScore = alternativeScores[i].TimbulanSampah + alternativeScores[i].JarakTpa + alternativeScores[i].KondisiTanah + alternativeScores[i].JarakPemukiman + alternativeScores[i].JarakSungai + alternativeScores[i].PartisipasiMasyarakat + alternativeScores[i].CakupanRumah + alternativeScores[i].Aksesibilitas
	}

	_, err = s.Repository.CreateFinalScore(ctx, finalscores)

	if err != nil {
		return nil, err
	}

	return finalscores, nil
}
