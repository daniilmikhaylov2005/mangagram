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

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{
			Error: err.Error(),
		})
	}

	user.Password = hashedPassword

	rUser, err := repository.InsertUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, rUser)
}

func Login(c echo.Context) error {
	var user models.LoginUser
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{
			Error: err.Error(),
		})
	}

	userFromDb, err := repository.SelectUserByEmail(user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{
			Error: err.Error(),
		})
	}

	if err := CheckPasswordHash(userFromDb.Password, user.Password); err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{
			Error: err.Error(),
		})
	}

	token, err := CreateToken(userFromDb.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Token{
		Token: token,
	})
}
