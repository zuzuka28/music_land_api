package model

type ReactionSearchQuery struct {
	UserID             string
	ReactionType       ReactionType
	ReactionTargetID   string
	ReactionTargetType ReactionTargetType
}
