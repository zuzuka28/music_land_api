package user

import "time"

type user struct {
	Username string `xorm:"not null unique"`
	Password string `xorm:"varchar(200)"`
	Salt     string
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}
