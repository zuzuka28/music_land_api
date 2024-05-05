package auth

import (
	"context"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type (
	Service interface {
		Authenticate(ctx context.Context, cmd *model.AuthCommand) (*model.Identity, error)
	}
)
