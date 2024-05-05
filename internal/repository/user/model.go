package user

import "time"

type user struct {
	ID       string
	Username string `xorm:"not null unique"`
	Password string `xorm:"varchar(200)"`
	Salt     string
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}
