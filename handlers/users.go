package handlers

import (
	"net/http"

	"github.com/daniilmikhaylov2005/mangagram/models"
	"github.com/daniilmikhaylov2005/mangagram/repository"
	"github.com/labstack/echo/v4"
)

func Signup(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{
			Error: err.Error(),
		})
	}

	user, err := repository.InsertUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	return nil
}
