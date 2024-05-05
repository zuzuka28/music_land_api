package user

import (
	"context"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type (
	Service interface {
		Create(ctx context.Context, cmd *model.UserCreateCommand) error
	}
)
