package auth

import (
	"github.com/labstack/echo/v4"
	dto "ta13-svc/internal/dto/auth"
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

// Login
// @Summary Login auth
// @Description Login auth
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.AuthLoginRequest true "request body"
// @Success 200 {object} dto.AuthLoginResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /auth/login [post]
func (h *handler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.AuthLoginRequest)

	if err = c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	if err = c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	data, err := h.service.Login(ctx, payload)

	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(data).Send(c)
}

// Register
// @Summary Register auth
// @Description Register auth
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.AuthRegisterRequest true "request body"
// @Success 200 {object} dto.AuthRegisterResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /auth/register [post]
func (h *handler) Register(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.AuthRegisterRequest)

	if err = c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	data, err := h.service.Register(ctx, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(data).Send(c)
}
