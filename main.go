package main

import (
	"example/web-service-gin/album"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize the album service
	albumService := album.NewInMemoryAlbumService()

	// Routes
	e.GET("/albums", func(c echo.Context) error {
		albums, err := albumService.GetAlbums()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, albums)
	})

	e.GET("/albums/:id", func(c echo.Context) error {
		id := c.Param("id")
		album, err := albumService.GetAlbumByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "album not found"})
		}
		return c.JSON(http.StatusOK, album)
	})

	e.POST("/albums", func(c echo.Context) error {
		var newAlbum album.Album
		if err := c.Bind(&newAlbum); err != nil {
			return err
		}
		if err := albumService.AddAlbum(newAlbum); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusCreated, newAlbum)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
