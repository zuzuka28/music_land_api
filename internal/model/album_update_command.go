package model

type AlbumAction string

type AlbumUpdateCommand struct {
	AlbumID  string
	TrackIDs []string
}
