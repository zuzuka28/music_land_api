package model

type ReactionType string

const (
	ReactionTypeLike    ReactionType = "like"
	ReactionTypeDislike ReactionType = "dislike"
)

type ReactionTargetType string

const (
	ReactionTargetTypeTrack ReactionTargetType = "track"
	ReactionTargetTypeAlbum ReactionTargetType = "album"
)

type Reaction struct {
	UserID             string
	ReactionType       ReactionType
	ReactionTargetID   string
	ReactionTargetType ReactionTargetType
}
