package user

type createRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type createResponse struct {
	Status string `json:"status"`
}
