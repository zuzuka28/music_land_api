package auth

import (
	"context"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type (
	UserService interface {
		Fetch(ctx context.Context, query *model.UserFetchQuery) (*model.User, error)
	}
)
