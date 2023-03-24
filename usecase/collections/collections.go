package collections

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ta13-svc/data/mysql"
	"ta13-svc/dto/collections"
	"ta13-svc/entity"
)

func GetRepoCollections() ([]entity.Collection, error) {
	db := mysql.GetDBInstance()
	var collections []entity.Collection

	if err := db.Find(&collections).Error; err != nil {
		return nil, err
	}

	return collections, nil
}

// GetCollections godoc
// @Summary Get all collections
// @Accept json
// @Produce json
// @Success 200 {object} ResponseGetCollections "Success"
// @Router /collections [get]
// @tags Collections
func GetCollections(c echo.Context) error {
	results, _ := GetRepoCollections()
	return c.JSON(http.StatusOK, collections.ResponseGetCollections{
		Collections: results,
	})
}
