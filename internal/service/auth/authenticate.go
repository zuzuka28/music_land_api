package auth

import (
	"context"
	"crypto/subtle"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
	"github.com/zuzuka28/music_land_api/pkg/passencode"
)

type userFetcher interface {
	Fetch(ctx context.Context, query *model.UserFetchQuery) (*model.User, error)
}

type authenticateService struct {
	r userFetcher
}

func newAuthenticateService(r userFetcher) *authenticateService {
	return &authenticateService{
		r: r,
	}
}

func (s *authenticateService) Authenticate(
	ctx context.Context,
	cmd *model.AuthCommand,
) (*model.Identity, error) {
	usr, err := s.r.Fetch(ctx, &model.UserFetchQuery{
		Nickname: cmd.Login,
	})
	if err != nil {
		return nil, fmt.Errorf("fetch user: %w", err)
	}

	if ok := comparePassword(cmd.Password, usr.Salt, usr.Password); !ok {
		return nil, fmt.Errorf("%w: bad password", model.ErrNotValid)
	}

	return &model.Identity{
		ID:       usr.ID,
		Nickname: usr.Username,
	}, nil
}

func comparePassword(password, salt, hash string) bool {
	hashedPassword, _ := passencode.EncodePassword(password, salt)
	return subtle.ConstantTimeCompare([]byte(hashedPassword), []byte(hash)) == 1
}
