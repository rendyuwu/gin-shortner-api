package web

type ShortenerUpdateRequest struct {
	ID         int    `json:"id,omitempty"`
	CustomCode string `json:"custom_code,omitempty" binding:"max=255" json:"custom_code,omitempty"`
}
