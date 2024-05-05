package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
	"github.com/zuzuka28/music_land_api/pkg/passencode"
)

type createService struct {
	r Repository
}

func newCreateService(r Repository) *createService {
	return &createService{
		r: r,
	}
}

func (s *createService) Create(ctx context.Context, cmd *model.UserCreateCommand) error {
	_, err := s.r.Fetch(ctx, cmd.Username)

	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return fmt.Errorf("check user exists: %w", err)
	}

	if err == nil {
		return fmt.Errorf("%w: user already exists", model.ErrNotValid)
	}

	salt, _ := passencode.RandomString(10) //nolint:gomnd
	password, _ := passencode.EncodePassword(string(cmd.Password), salt)

	usr := &model.User{
		ID:       "",
		Username: cmd.Username,
		Password: password,
		Salt:     salt,
	}

	if err := s.r.Create(ctx, usr); err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}
