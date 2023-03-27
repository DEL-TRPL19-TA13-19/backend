package collection

import (
	"context"
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {

	g.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// DO PROCESS jwt decode
			// GET USER ID then set
			ctx := context.WithValue(c.Request().Context(), "user", "9f9b019f-5fd1-4383-bf71-7d43f7b10f9e")
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	})
	g.GET("/:user_id", h.GetByUserID)
	g.POST("", h.Create)
	g.PATCH("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}
