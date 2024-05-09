package model

type AlbumAction string

type AlbumUpdateCommand struct {
	OwnerID  string
	AlbumID  string
	TrackIDs []string
}
