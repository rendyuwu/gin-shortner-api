package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rendyuwu/gin-shortner-api/domain/web"
	"github.com/rendyuwu/gin-shortner-api/exception"
	"github.com/rendyuwu/gin-shortner-api/helper"
	"github.com/rendyuwu/gin-shortner-api/service"
	"net/http"
	"strconv"
)

type ShortenerController interface {
	FindById(c *gin.Context)
	FindByCode(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type ShortenerControllerImpl struct {
	ShortenerService service.ShortenerService
}

func NewShortenerController(shortenerService service.ShortenerService) ShortenerController {
	return &ShortenerControllerImpl{
		ShortenerService: shortenerService,
	}
}

func (controller ShortenerControllerImpl) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idShortener"))
	helper.PanicIfError(err)

	shortener := controller.ShortenerService.FindById(id)

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   shortener,
	})
}
func (controller ShortenerControllerImpl) FindByCode(c *gin.Context) {
	shortener := controller.ShortenerService.FindByCode(c.Param("idShortener"))

	c.Redirect(http.StatusMovedPermanently, shortener.Url)
}

func (controller ShortenerControllerImpl) Create(c *gin.Context) {
	var shortenerRequest = web.ShortenerCreateRequest{}

	err := c.ShouldBindJSON(&shortenerRequest)
	if exception.ValidationError(c, err) {
		return
	}

	shortener := controller.ShortenerService.Create(shortenerRequest)

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   shortener,
	})
}

func (controller ShortenerControllerImpl) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idShortener"))
	helper.PanicIfError(err)

	shortenerRequest := web.ShortenerUpdateRequest{
		ID: id,
	}
	err = c.ShouldBindJSON(&shortenerRequest)
	exception.ValidationError(c, err)

	shortener := controller.ShortenerService.Update(shortenerRequest)

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   shortener,
	})
}

func (controller ShortenerControllerImpl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("bookId"))
	helper.PanicIfError(err)

	controller.ShortenerService.Delete(id)

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	})
}
