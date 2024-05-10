package track

import "time"

type track struct {
	UID     string
	Name    string
	Author  string
	FileID  string
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
