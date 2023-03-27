package alternative

import (
	"fmt"
	"github.com/labstack/echo/v4"
	dto "ta13-svc/internal/dto/alternative"
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

// GetAll
// @Summary Get All Alternatives
// @Description Get All Alternatives
// @Tags alternative
// @Accept json
// @Produce json
// @Success 200 {object} dto.AlternativesGetResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /alternative [get]
func (h *handler) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := h.service.FindAll(ctx)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// GetByCollectionID
// @Summary Get Alternatives By Collection ID
// @Description Get Alternatives By Collection ID
// @Tags alternative
// @Accept json
// @Produce json
// @Param collection_id path string true "id path"
// @Success 200 {object} dto.AlternativeGetByCollectionIDResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /alternative/{collection_id} [GET]
func (h *handler) GetByCollectionID(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.AlternativeGetByCollectionIDRequest)
	if err = c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		response := response.ErrorBuilder(&response.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	fmt.Printf("IDCOLLECTION : %+v", payload)

	result, err := h.service.FindByCollectionID(ctx, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// Create godoc
// @Summary Create Alternative
// @Description Create Alternative
// @Tags alternative
// @Accept  json
// @Produce  json
// @Param request body dto.AlternativeCreateRequest true "request body"
// @Success 200 {object} dto.AlternativeCreateResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /alternative [post]
func (h *handler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.AlternativeCreateRequest)
	fmt.Printf("PAYLOAD CREATECOLELCTION : %+v", payload)

	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Create(ctx, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// Update godoc
// @Summary Update alternative
// @Description Update alternative
// @Tags alternative
// @Accept  json
// @Produce  json
// @Param id path string true "id path"
// @Param request body dto.AlternativeUpdateRequest true "request body"
// @Success 200 {object} dto.AlternativeUpdateResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /alternative/{id} [patch]
func (h *handler) Update(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.AlternativeUpdateRequest)
	if err := c.Bind(&payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Update(ctx, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// Delete godoc
// @Summary Delete alternative
// @Description Delete alternative
// @Tags alternative
// @Accept  json
// @Produce  json
// @Param id path string true "id path"
// @Success 200 {object}  dto.AlternativeDeleteResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /alternative/{id} [delete]
func (h *handler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.AlternativeDeleteRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Delete(ctx, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}
