package reaction

type searchRequest struct {
	UserID             string `form:"userID"`
	ReactionType       string `form:"reactionType"`
	ReactionTargetID   string `form:"reactionTargetID"`
	ReactionTargetType string `form:"reactionTargetType"`
}

type reaction struct {
	UserID             string `json:"userID"`
	ReactionType       string `json:"reactionType"`
	ReactionTargetID   string `json:"reactionTargetID"`
	ReactionTargetType string `json:"reactionTargetType"`
}

type searchResponse struct {
	Items []*reaction `json:"items"`
}

type createRequest struct {
	ReactionType       string `form:"reactionType" json:"reactionType" binding:"required"`
	ReactionTargetID   string `form:"reactionTargetID" json:"reactionTargetID" binding:"required"`
	ReactionTargetType string `form:"reactionTargetType" json:"reactionTargetType" binding:"required"`
}

type createResponse struct {
	Status string `json:"status"`
}

type deleteRequest struct {
	ReactionTargetID string `form:"reactionTargetID" json:"reactionTargetID" binding:"required"`
}

type deleteResponse struct {
	Status string `json:"status"`
}
