package model

type AlbumCreateCommand struct {
	OwnerID  string
	Name     string
	TrackIDs []string
}
