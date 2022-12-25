package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/daniilmikhaylov2005/mangagram/models"
	"github.com/daniilmikhaylov2005/mangagram/repository"
	"github.com/labstack/echo/v4"
)

func GetAllImages(c echo.Context) error {
	images, err := repository.SelectAllImages()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{
			Error: "Something went wrong on server",
		})
	}

	return c.JSON(http.StatusOK, images)
}

func GetImageById(c echo.Context) error {
	stringId := c.Param("id")
	intId, err := strconv.Atoi(stringId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{
			Error: "id must be a numeric",
		})
	}

	image, err := repository.SelectImageById(intId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{
			Error: "Something went wrong on server",
		})
	}

	return c.JSON(http.StatusOK, image)
}

func UploadImage(c echo.Context) error {
	// Src
	file, err := c.FormFile("image")

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{
			Error: err.Error(),
		})
	}

	src, err := file.Open()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{
			Error: err.Error(),
		})
	}

	defer src.Close()

	// unique filename
	words := strings.Split(file.Filename, ".")
	v, _ := time.Now().UTC().MarshalText()
	words[0] += string(v)
	newFilename := strings.Join(words, ".")

	// Destination
	dst, err := os.Create("assets/images/" + newFilename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{
			Error: err.Error(),
		})
	}

	defer dst.Close()

	var image models.Image
	image.Url = fmt.Sprintf("assets/images/%s", newFilename)
	image.Author = "User"

	err = repository.InsertImage(image)
	if err != nil {
		e := os.Remove(image.Url)
		if e != nil {
			return c.JSON(http.StatusInternalServerError, models.Error{
				Error: e.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, models.Error{
			Error: err.Error(),
		})
	}

	// Copy
	if _, err := io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.Status{
		Status: "Image Uploaded",
	})
}
