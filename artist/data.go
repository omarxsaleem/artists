package artist

import "errors"

// InMemoryArtistService is a concrete implementation of ArtistService
// that stores artists in memory.
type InMemoryArtistService struct {
	artists []Artist
}

// NewInMemoryArtistService initializes an InMemoryArtistService with some data.
func NewInMemoryArtistService() *InMemoryArtistService {
	return &InMemoryArtistService{
		artists: []Artist{
			{ID: "1", Name: "John Coltrane"},
			{ID: "2", Name: "Gerry Mulligan"},
			{ID: "3", Name: "Sarah Vaughan"},
		},
	}
}

// GetArtists returns a list of all artists.
func (s *InMemoryArtistService) GetArtists() ([]Artist, error) {
	return s.artists, nil
}

// GetArtistByID returns an artist by its ID.
func (s *InMemoryArtistService) GetArtistByID(id string) (Artist, error) {
	for _, a := range s.artists {
		if a.ID == id {
			return a, nil
		}
	}
	return Artist{}, errors.New("artist not found")
}

// AddArtist adds a new artist to the collection.
func (s *InMemoryArtistService) AddArtist(artist Artist) error {
	s.artists = append(s.artists, artist)
	return nil
}
