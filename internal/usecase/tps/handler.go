package tps

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"ta13-svc/internal/dto/tps"
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

// Get
// @Summary Get All Tps
// @Description Get All Tps
// @Tags tps
// @Accept json
// @Produce json
// @Success 200 {object} dto.TpsGetResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /tps [get]
func (h *handler) Get(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := h.service.FindAll(ctx)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// GetByID
// @Summary Get Tps by id
// @Description Get Tps by id
// @Tags tps
// @Accept json
// @Produce json
// @Param id path string true "id path"
// @Success 200 {object} dto.TpsGetByIdResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /tps/{id} [get]
func (h *handler) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.TpsGetByIdRequest)

	if err = c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		response := response.ErrorBuilder(&response.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	fmt.Printf("%+v", payload)

	result, err := h.service.FindById(ctx, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// Create godoc
// @Summary Create tps
// @Description Create tps
// @Tags tps
// @Accept  json
// @Produce  json
// @Param request body dto.TpsCreateRequest true "request body"
// @Success 200 {object} dto.TpsCreateResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /tps [post]
func (h *handler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	payload := new(dto.TpsCreateRequest)

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
// @Summary Update tps
// @Description Update tps
// @Tags tps
// @Accept  json
// @Produce  json
// @Param request body dto.TpsUpdateRequest true "request body"
// @Success 200 {object} dto.TpsUpdateResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /tps/{id} [patch]
func (h *handler) Update(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.TpsUpdateRequest)
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
// @Summary Delete tps
// @Description Delete tps
// @Tags tps
// @Accept  json
// @Produce  json
// @Param id path string true "id path"
// @Success 200 {object}  dto.TpsDeleteResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /tps/{id} [delete]
func (h *handler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.TpsDeleteRequest)
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
