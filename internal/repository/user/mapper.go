package user

import (
	"time"

	"github.com/zuzuka28/music_land_api/internal/model"
)

func mapUserToModel(in *user) *model.User {
	if in == nil {
		return nil
	}

	return &model.User{
		Username: in.Username,
		Password: in.Password,
		Salt:     in.Salt,
	}
}

func mapUserToInternal(in *model.User) *user {
	if in == nil {
		return nil
	}

	return &user{
		Username: in.Username,
		Password: in.Password,
		Salt:     in.Salt,
		Created:  time.Time{},
		Updated:  time.Time{},
	}
}
