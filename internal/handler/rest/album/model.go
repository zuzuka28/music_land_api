package album

type fetchResponse struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	OwnerID  string   `json:"ownerID"`
	TrackIDs []string `json:"tracks"`
}

type createRequest struct {
	Name     string   `form:"name" json:"name" binding:"required"`
	TrackIDs []string `form:"tracks" json:"tracks" binding:"required"`
}

type createResponse struct {
	Status string `json:"status"`
}

type deleteResponse struct {
	Status string `json:"status"`
}

type updateRequest struct {
	TrackIDs []string `form:"tracks" json:"tracks" binding:"required"`
}

type updateResponse struct {
	Status string `json:"status"`
}
