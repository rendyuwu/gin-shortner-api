package web

type ShortenerUpdateRequest struct {
	Id         int    `json:"id,omitempty"`
	CustomCode string `json:"custom_code,omitempty" binding:"max=255" json:"custom_code,omitempty"`
	Url        string `json:"url,omitempty" binding:"required,max=1000,uri" json:"url,omitempty"`
}
