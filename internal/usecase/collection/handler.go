package collection

import (
	"fmt"
	"github.com/labstack/echo/v4"
	dto "ta13-svc/internal/dto/collection"
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

// GetByUserID
// @Summary Get Collection By UserID
// @Description Get Collection By UserID
// @Tags collection
// @Accept json
// @Produce json
// @Param user_id path string true "id path"
// @Success 200 {object} dto.CollectionGetByUserIDResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /collection/{user_id} [get]
func (h *handler) GetByUserID(c echo.Context) error {
	ctx := c.Request().Context()
	payload := new(dto.CollectionGetByUserIDRequest)
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
// @Summary Create collection
// @Description Create collection
// @Tags collection
// @Accept  json
// @Produce  json
// @Param request body dto.CollectionCreateRequest true "request body"
// @Success 200 {object} dto.CollectionCreateResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /collection [post]
func (h *handler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	payload := new(dto.CollectionCreateRequest)

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
// @Summary Update collection
// @Description Update collection
// @Tags collection
// @Accept  json
// @Produce  json
// @Param id path string true "id path"
// @Param request body dto.CollectionUpdateRequest true "request body"
// @Success 200 {object} dto.CollectionUpdateResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /collection/{id} [patch]
func (h *handler) Update(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.CollectionUpdateRequest)
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
// @Summary Delete collection
// @Description Delete collection
// @Tags collection
// @Accept  json
// @Produce  json
// @Param id path string true "id path"
// @Success 200 {object}  dto.CollectionDeleteResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /collection/{id} [delete]
func (h *handler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.CollectionDeleteRequest)
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
