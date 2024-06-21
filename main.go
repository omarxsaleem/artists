package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// album represents data about a record album.
type artist struct {
	ID   string `json: "id"`
	Name string `json: "name"`
}

// albums slice to seed record album data.
var artists = []artist{
	{ID: "1", Name: "John Coltrane"},
	{ID: "2", Name: "Gerry Mulligan"},
	{ID: "3", Name: "Sarah Vaughan"},
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/artists", getArtists)
	e.GET("/artists/:id", getAlbumByID)
	e.POST("/artists", postArtists)

	e.Logger.Fatal(e.Start(":8080"))
}

// getAlbums responds with the list of all albums as JSON.
func getArtists(c echo.Context) error {
	return c.JSON(http.StatusOK, artists)
}

// postAlbums adds an album from JSON received in the request body.
func postArtists(c echo.Context) error {
	var newArtist artist

	if err := c.Bind(&newArtist); err != nil {
		return err
	}

	artists = append(artists, newArtist)
	return c.JSON(http.StatusCreated, newArtist)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c echo.Context) error {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range artists {
		if a.ID == id {
			return c.JSON(http.StatusOK, a)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "album not found"})
}
