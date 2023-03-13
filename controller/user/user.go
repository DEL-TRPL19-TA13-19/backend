package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ta13-svc/model"
	"ta13-svc/storage"
)

func GetRepoStudents() ([]model.User, error) {
	db := storage.GetDBInstance()
	var users []model.User

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// GetUsers godoc
// @Summary Get all users
// @Accept json
// @Produce json
// @Success 200 {object} ResponseGetUsers "Success"
// @Router /users [get]
// @tags User
func GetUsers(c echo.Context) error {
	users, _ := GetRepoStudents()
	return c.JSON(http.StatusOK, ResponseGetUsers{Users: users})
}

// GetHello godoc
// @Summary Get hi
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHello "Success"
// @Router / [get]
// @tags hai
func GetHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hai")
}
