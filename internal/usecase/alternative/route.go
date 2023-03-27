package alternative

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.GetAll)
	g.GET("/:collection_id", h.GetByCollectionID)
	g.POST("", h.Create)
	g.PATCH("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}
