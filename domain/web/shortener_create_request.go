package web

type ShortenerCreateRequest struct {
	CustomCode string `json:"custom_code,omitempty" binding:"max=255"`
	Url        string `json:"url,omitempty" binding:"required,max=1000,uri"`
}
