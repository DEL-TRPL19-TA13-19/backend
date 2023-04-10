package ahp

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	g.GET("/criteria", h.GetCriteria)
	g.GET("/scores/:collection_id", h.GetScores)
	g.GET("/finalscores/:collection_id", h.GetFinalScores)
}
