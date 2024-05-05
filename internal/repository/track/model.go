package track

import "time"

type track struct {
	ID      int64 `xorm:"autoincr"`
	Name    string
	Author  string
	FileID  string
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
