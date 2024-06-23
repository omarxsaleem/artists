package album

// AlbumService defines the methods that any
// service handling albums must implement.
type AlbumService interface {
	GetAlbums() ([]Album, error)
	GetAlbumByID(id string) (Album, error)
	AddAlbum(album Album) error
}
