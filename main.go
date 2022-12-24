package main

import (
	"net/http"

	"github.com/daniilmikhaylov2005/mangagram/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/assets/images/:url", func(c echo.Context) error {
		url := c.Param("url")
		return c.File("assets/images/" + url)
	})

	auth := e.Group("/auth")
	auth.POST("/signup", handlers.Signup)
	auth.POST("/signin", handlers.Login)

	images := e.Group("/images")
	images.POST("/", handlers.UploadImage)
	images.GET("/", handlers.GetAllImages)
	images.GET("/:id", handlers.GetImageById)

	e.Logger.Fatal(e.Start(":8080"))
}
