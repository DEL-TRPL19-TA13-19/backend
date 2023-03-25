package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ta13-svc/pkg/response"
)

func ErrorHandler(err error, c echo.Context) {
	var errCustom *response.Error

	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	switch report.Code {
	case http.StatusNotFound:
		errCustom = response.ErrorBuilder(&response.ErrorConstant.RouteNotFound, err)
	default:
		errCustom = response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	response.ErrorResponse(errCustom).Send(c)
}
