package model

import (
	"time"
)

type Shortener struct {
	ID         int
	Code       string
	CustomCode string
	Url        string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
