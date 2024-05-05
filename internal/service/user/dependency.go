package user

import (
	"context"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type (
	Repository interface {
		Fetch(ctx context.Context, id string) (*model.User, error)
		Create(ctx context.Context, item *model.User) error
	}
)
