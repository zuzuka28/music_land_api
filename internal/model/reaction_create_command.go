package model

type ReactionCreateCommand struct {
	UserID             string
	ReactionType       ReactionType
	ReactionTargetID   string
	ReactionTargetType ReactionTargetType
}
