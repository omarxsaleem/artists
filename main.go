package main

import (
	"example/web-service-gin/album"
	"example/web-service-gin/artist"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize the services
	albumService := album.NewInMemoryAlbumService()
	artistService := artist.NewInMemoryArtistService()

	// Album routes
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

	// Artist routes
	e.GET("/artists", func(c echo.Context) error {
		artists, err := artistService.GetArtists()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, artists)
	})

	e.GET("/artists/:id", func(c echo.Context) error {
		id := c.Param("id")
		artist, err := artistService.GetArtistByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "artist not found"})
		}
		return c.JSON(http.StatusOK, artist)
	})

	e.POST("/artists", func(c echo.Context) error {
		var newArtist artist.Artist
		if err := c.Bind(&newArtist); err != nil {
			return err
		}
		if err := artistService.AddArtist(newArtist); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusCreated, newArtist)
	})

	// Get all albums by artist ID
	e.GET("/albums/artist/:artist_id", func(c echo.Context) error {
		artistID := c.Param("artist_id")
		allAlbums, err := albumService.GetAlbums()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		var albumsByArtist []album.Album
		for _, a := range allAlbums {
			if a.ArtistID == artistID {
				albumsByArtist = append(albumsByArtist, a)
			}
		}

		return c.JSON(http.StatusOK, albumsByArtist)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
