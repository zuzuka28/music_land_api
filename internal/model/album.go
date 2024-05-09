package model

type Album struct {
	ID       string
	Name     string
	OwnerID  string
	TrackIDs []string
}
