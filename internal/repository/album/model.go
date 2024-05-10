package album

import "time"

type album struct {
	UID     string
	Name    string
	OwnerID string
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type albumItem struct {
	AlbumID string
	TrackID string
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
