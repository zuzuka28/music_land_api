package reaction

import (
	"time"

	"github.com/zuzuka28/music_land_api/internal/model"
)

func mapReactionToInternal(in *model.Reaction) *reaction {
	if in == nil {
		return nil
	}

	return &reaction{
		UserID:             in.UserID,
		ReactionType:       string(in.ReactionType),
		ReactionTargetID:   in.ReactionTargetID,
		ReactionTargetType: string(in.ReactionTargetType),
		Created:            time.Time{},
		Updated:            time.Time{},
	}
}

func mapReactionToModel(in *reaction) *model.Reaction {
	if in == nil {
		return nil
	}

	return &model.Reaction{
		UserID:             in.UserID,
		ReactionType:       model.ReactionType(in.ReactionType),
		ReactionTargetID:   in.ReactionTargetID,
		ReactionTargetType: model.ReactionTargetType(in.ReactionTargetType),
	}
}
