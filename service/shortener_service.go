package service

import (
	"github.com/rendyuwu/gin-shortner-api/domain/model"
	"github.com/rendyuwu/gin-shortner-api/domain/web"
	"github.com/rendyuwu/gin-shortner-api/helper"
	"github.com/rendyuwu/gin-shortner-api/repository"
)

type ShortenerService interface {
	FindAll() []web.ShortenerResponse
	FindById(id int) web.ShortenerResponse
	Create(request web.ShortenerCreateRequest) web.ShortenerResponse
	Update(request web.ShortenerUpdateRequest) web.ShortenerResponse
	Delete(id int)
}
type ShortenerServiceImpl struct {
	ShortenerRepository repository.ShortenerRepository
}

func (service ShortenerServiceImpl) FindAll() []web.ShortenerResponse {
	shortener, err := service.ShortenerRepository.FindAll()
	helper.IsNotFoundError(err)

	return helper.ToShortenerResponses(shortener)
}

func (service ShortenerServiceImpl) FindById(id int) web.ShortenerResponse {
	shortener, err := service.ShortenerRepository.FindById(id)
	helper.IsNotFoundError(err)

	return helper.ToShortenerResponse(shortener)
}

func (service ShortenerServiceImpl) Create(request web.ShortenerCreateRequest) web.ShortenerResponse {
	// generate random 6 string
	var code string

	// check if code already exist
	for {
		code = helper.RandomString(6)
		_, err := service.ShortenerRepository.FindByCode(code)
		if err != nil {
			continue
		} else {
			break
		}
	}

	shortenerRequest := model.Shortener{
		Code:       code,
		CustomCode: request.CustomCode,
		Url:        request.Url,
	}

	shortener, err := service.ShortenerRepository.Create(shortenerRequest)
	helper.PanicIfError(err)

	return helper.ToShortenerResponse(shortener)
}

func (service ShortenerServiceImpl) Update(request web.ShortenerUpdateRequest) web.ShortenerResponse {
	shortener, err := service.ShortenerRepository.FindById(request.ID)
	helper.IsNotFoundError(err)

	shortener.CustomCode = request.CustomCode

	newShortener, err := service.ShortenerRepository.Update(shortener)

	return helper.ToShortenerResponse(newShortener)
}

func (service ShortenerServiceImpl) Delete(id int) {
	shortener, err := service.ShortenerRepository.FindById(id)
	helper.IsNotFoundError(err)

	err = service.ShortenerRepository.Delete(shortener)
	helper.PanicIfError(err)
}
