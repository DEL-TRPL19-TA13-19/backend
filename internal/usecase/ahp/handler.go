package ahp

import (
	"github.com/labstack/echo/v4"
	dto "ta13-svc/internal/dto/ahp"
	"ta13-svc/internal/factory"
	"ta13-svc/pkg/response"
)

type handler struct {
	service *service
}

var err error

func NewHandler(f *factory.Factory) *handler {
	service := NewService(f)
	return &handler{service}
}

// GetCriteria
// @Summary Get All Criteria Alternative
// @Description Get ALl Criteria Alternative
// @Tags AHP
// @Accept json
// @Produce json
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /ahp/criteria [get]
func (h *handler) GetCriteria(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := h.service.FindCriteriaAlternative(ctx)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// GetScores
// @Summary Get Scores By Collection ID
// @Description Get Scores By Collection ID
// @Tags AHP
// @Accept json
// @Produce json
// @Param collection_id path string true "collection_id path"
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /ahp/scores/{collection_id} [get]
func (h *handler) GetScores(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.AHPByCollectionIDRequest)
	if err = c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		response := response.ErrorBuilder(&response.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	result, err := h.service.FindScoreByCollectionID(ctx, &payload.CollectionID)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// GetFinalScores
// @Summary Get Final Scores
// @Description Get Final Scores
// @Tags AHP
// @Accept json
// @Produce json
// @Param collection_id path string true "collection_id path"
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /ahp/final_scores/{collection_id} [get]
func (h *handler) GetFinalScores(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.AHPByCollectionIDRequest)
	if err = c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		response := response.ErrorBuilder(&response.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	result, err := h.service.FindFinalScoreByCollectionID(ctx, &payload.CollectionID)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// CalculateScores
// @Summary Calculate Scores by Collection ID
// @Description Calculate Scores by Collection ID
// @Tags AHP
// @Accept json
// @Produce json
// @Param collection_id path string true "collection_id path"
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /ahp/scores/calculate/{collection_id} [get]
func (h *handler) CalculateScores(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.AHPByCollectionIDRequest)
	if err = c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		response := response.ErrorBuilder(&response.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	result, err := h.service.CalculateScoreAlternativeByCollectionID(ctx, &payload.CollectionID)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// CalculateFinalScores
// @Summary Calculate Final Scores by Collection ID
// @Description Calculate Final Scores by Collection ID
// @Tags AHP
// @Accept json
// @Produce json
// @Param collection_id path string true "collection_id path"
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /ahp/final_scores/calculate/{collection_id} [get]
func (h *handler) CalculateFinalScores(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.AHPByCollectionIDRequest)
	if err = c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		response := response.ErrorBuilder(&response.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	result, err := h.service.CalculateFinalScoreByCollectionID(ctx, &payload.CollectionID)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}
