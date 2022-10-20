package web

type ShortenerResponse struct {
	Id         int    `json:"id,omitempty"`
	Code       string `json:"code,omitempty"`
	CustomCode string `json:"custom_code,omitempty"`
	Url        string `json:"url,omitempty"`
}
