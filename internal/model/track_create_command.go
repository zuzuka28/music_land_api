package model

type TrackCreateCommand struct {
	Name      string
	Author    string
	TrackData []byte
}
