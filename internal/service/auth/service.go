package auth

import (
	"context"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type Service struct {
	*authenticateService
}

func NewService(us UserService) *Service {
	return &Service{
		authenticateService: newAuthenticateService(us),
	}
}

func (s *Service) Authenticate(ctx context.Context, cmd *model.AuthCommand) (*model.Identity, error) {
	return s.authenticateService.Authenticate(ctx, cmd)
}
