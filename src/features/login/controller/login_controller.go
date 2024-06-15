package controller

import (
	"net/http"
	"academy/src/features/login/model"
	"github.com/labstack/echo/v4"
)

func RegisterUser(c echo.Context) error {
	var userRequest model.UserRequest
	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, Error{})
	}
}