package artist

// Artist represents data about a musical artist.
type Artist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ArtistService defines the methods that any
// service handling artists must implement.
type ArtistService interface {
	GetArtists() ([]Artist, error)
	GetArtistByID(id string) (Artist, error)
	AddArtist(artist Artist) error
}
