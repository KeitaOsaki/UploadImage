package view


type ImagesUrlResponse struct {
	Urls []string `json:"urls"`
}

func UploadResponse(urls []string)ImagesUrlResponse {
	return ImagesUrlResponse{
		Urls: urls,
	}
}