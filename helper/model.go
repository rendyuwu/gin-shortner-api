package helper

import (
	"github.com/rendyuwu/gin-shortner-api/domain/model"
	"github.com/rendyuwu/gin-shortner-api/domain/web"
)

func ToShortenerResponse(shortener model.Shortener) web.ShortenerResponse {
	return web.ShortenerResponse{
		ID:         shortener.ID,
		Code:       shortener.Code,
		CustomCode: shortener.CustomCode,
		Url:        shortener.Url,
	}
}

func ToShortenerResponses(shorteners []model.Shortener) []web.ShortenerResponse {
	var shortenerResponses []web.ShortenerResponse

	for _, shortener := range shorteners {
		shortenerResponse := web.ShortenerResponse{
			ID:         shortener.ID,
			Code:       shortener.Code,
			CustomCode: shortener.CustomCode,
			Url:        shortener.Url,
		}
		shortenerResponses = append(shortenerResponses, shortenerResponse)
	}
	return shortenerResponses
}
