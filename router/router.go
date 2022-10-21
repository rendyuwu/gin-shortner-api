package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rendyuwu/gin-shortner-api/controller"
	"github.com/rendyuwu/gin-shortner-api/exception"
)

func NewRouter(controller controller.ShortenerController) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), exception.Recovery())

	//route root
	router.GET("/:idShortener", controller.FindByCode)

	// route api
	v1 := router.Group("v1")

	//v1.GET("/books", controller.FindAll)
	v1.GET("/shortener/:idShortener", controller.FindById)
	v1.POST("/shortener", controller.Create)
	v1.PUT("/shortener/:idShortener", controller.Update)
	v1.DELETE("/shortener/:idShortener", controller.Delete)

	return router
}
