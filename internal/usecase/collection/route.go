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
			ctx := context.WithValue(c.Request().Context(), "user", "d64b33d0-62b7-413e-bdf9-c91079ffaf9a")
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	})
	g.GET("", h.Get)
	g.GET("/:id", h.GetByID)
	g.GET("/user/:user_id", h.GetByUserID)
	g.POST("", h.Create)
	g.PATCH("", h.Update)
	g.DELETE("/:id", h.Delete)
}
