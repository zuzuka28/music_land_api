package auth

import (
	"context"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type Service struct {
	tr Tracer

	as *authenticateService
}

func NewService(us UserService, tr Tracer) *Service {
	return &Service{
		tr: tr,
		as: newAuthenticateService(us),
	}
}

func (s *Service) Authenticate(ctx context.Context, cmd *model.AuthCommand) (*model.Identity, error) {
	ctx, span := s.tr.Start(ctx, "Authenticate")
	defer span.End()

	return s.as.Authenticate(ctx, cmd)
}
