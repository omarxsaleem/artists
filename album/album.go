package album

// Album represents data about a record album.
type Album struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	ArtistID string  `json:"artist"`
	Price    float64 `json:"price"`
}

// NewAlbum creates a new Album instance.
func NewAlbum(id, title, artist string, price float64) Album {
	return Album{
		ID:       id,
		Title:    title,
		ArtistID: artist,
		Price:    price,
	}
}
