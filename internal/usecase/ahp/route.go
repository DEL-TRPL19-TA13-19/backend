package ahp

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	g.GET("/criteria", h.GetCriteria)
	g.GET("/scores/:collection_id", h.GetScores)
	g.GET("/final_scores/:collection_id", h.GetFinalScores)
	g.GET("/scores/calculate/:collection_id", h.CalculateScores)
	g.GET("/final_scores/calculate/:collection_id", h.CalculateFinalScores)
}
