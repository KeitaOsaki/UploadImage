package model

type Image struct {
	Image string `json:"image"`
}

type UploadRequest struct {
	Images []Image `json:"images"`
}