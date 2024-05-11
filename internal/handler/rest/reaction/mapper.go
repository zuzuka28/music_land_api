package reaction

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zuzuka28/music_land_api/internal/handler/rest/middleware/auth"
	"github.com/zuzuka28/music_land_api/internal/model"
)

func parseReactionDeleteCommand(gctx *gin.Context) (*model.ReactionDeleteCommand, error) {
	req := new(deleteRequest)
	if err := gctx.ShouldBind(req); err != nil {
		return nil, fmt.Errorf("bind delete request: %w", err)
	}

	usr, err := parseUserInfo(gctx)
	if err != nil {
		return nil, fmt.Errorf("parse user info: %w", err)
	}

	return &model.ReactionDeleteCommand{
		UserID:           usr.ID,
		ReactionTargetID: req.ReactionTargetID,
	}, nil
}

func parseReactionCreateCommand(gctx *gin.Context) (*model.ReactionCreateCommand, error) {
	req := new(createRequest)
	if err := gctx.ShouldBind(req); err != nil {
		return nil, fmt.Errorf("bind create request: %w", err)
	}

	usr, err := parseUserInfo(gctx)
	if err != nil {
		return nil, fmt.Errorf("parse user info: %w", err)
	}

	// TODO: add action validation
	return &model.ReactionCreateCommand{
		UserID:             usr.ID,
		ReactionType:       model.ReactionType(req.ReactionType),
		ReactionTargetID:   req.ReactionTargetID,
		ReactionTargetType: model.ReactionTargetType(req.ReactionTargetType),
	}, nil
}

func parseReactionSearchQuery(gctx *gin.Context) (*model.ReactionSearchQuery, error) {
	req := new(searchRequest)
	if err := gctx.ShouldBind(req); err != nil {
		return nil, fmt.Errorf("bind search request: %w", err)
	}

	// TODO: add action validation
	return &model.ReactionSearchQuery{
		UserID:             req.UserID,
		ReactionType:       model.ReactionType(req.ReactionType),
		ReactionTargetID:   req.ReactionTargetID,
		ReactionTargetType: model.ReactionTargetType(req.ReactionTargetID),
	}, nil
}

func mapSearchResponse(in []*model.Reaction) *searchResponse {
	items := make([]*reaction, 0, len(in))

	for i := range in {
		items = append(items, mapReactionToResponse(in[i]))
	}

	return &searchResponse{
		Items: items,
	}
}

func mapReactionToResponse(in *model.Reaction) *reaction {
	if in == nil {
		return nil
	}

	return &reaction{
		UserID:             in.UserID,
		ReactionType:       string(in.ReactionType),
		ReactionTargetID:   in.ReactionTargetID,
		ReactionTargetType: string(in.ReactionTargetType),
	}
}

func parseUserInfo(gctx *gin.Context) (*model.Identity, error) {
	rawusr, has := gctx.Get(auth.IdentityKey())
	if !has {
		return nil, fmt.Errorf("%w: no user info", model.ErrNotValid)
	}

	usr, ok := rawusr.(*model.Identity)
	if !ok {
		return nil, fmt.Errorf("%w: invalid user info", model.ErrNotValid)
	}

	return usr, nil
}
