package album

import "errors"

// InMemoryAlbumService is a concrete implementation of AlbumService
// that stores albums in memory.
type InMemoryAlbumService struct {
	albums []Album
}

// NewInMemoryAlbumService initializes an InMemoryAlbumService with some data.
func NewInMemoryAlbumService() *InMemoryAlbumService {
	return &InMemoryAlbumService{
		albums: []Album{
			{ID: "1", Title: "Blue Train", ArtistID: "John Coltrane", Price: 56.99},
			{ID: "2", Title: "Jeru", ArtistID: "Gerry Mulligan", Price: 17.99},
			{ID: "3", Title: "Sarah Vaughan and Clifford Brown", ArtistID: "Sarah Vaughan", Price: 39.99},
		},
	}
}

// GetAlbums returns a list of all albums.
func (s *InMemoryAlbumService) GetAlbums() ([]Album, error) {
	return s.albums, nil
}

// GetAlbumByID returns an album by its ID.
func (s *InMemoryAlbumService) GetAlbumByID(id string) (Album, error) {
	for _, a := range s.albums {
		if a.ID == id {
			return a, nil
		}
	}
	return Album{}, errors.New("album not found")
}

// AddAlbum adds a new album to the collection.
func (s *InMemoryAlbumService) AddAlbum(album Album) error {
	s.albums = append(s.albums, album)
	return nil
}
