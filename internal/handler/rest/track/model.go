package track

import "mime/multipart"

type fetchResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	FileID string `json:"fileID"`
}

type searchRequest struct {
	Name string `form:"name"`
}

type track struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

type searchResponse struct {
	Items []*track `json:"items"`
}

type saveRequest struct {
	Name       string                `form:"name" json:"name" binding:"required"`
	Author     string                `form:"author" json:"author" binding:"required"`
	Attachment *multipart.FileHeader `form:"attachment" json:"attachment" binding:"required"`
}

type saveResponse struct {
	Status string `json:"status"`
}

type deleteResponse struct {
	Status string `json:"status"`
}
