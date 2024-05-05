package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zuzuka28/music_land_api/internal/model"
)

func parseUserCreateCommand(gctx *gin.Context) (*model.UserCreateCommand, error) {
	req := new(createRequest)
	if err := gctx.ShouldBind(req); err != nil {
		return nil, fmt.Errorf("bind create request: %w", err)
	}

	return &model.UserCreateCommand{
		Username: req.Username,
		Password: model.Password(req.Password),
	}, nil
}
