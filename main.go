package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rendyuwu/gin-shortner-api/app"
	"github.com/rendyuwu/gin-shortner-api/controller"
	"github.com/rendyuwu/gin-shortner-api/domain/model"
	"github.com/rendyuwu/gin-shortner-api/env"
	"github.com/rendyuwu/gin-shortner-api/helper"
	"github.com/rendyuwu/gin-shortner-api/repository"
	"github.com/rendyuwu/gin-shortner-api/router"
	"github.com/rendyuwu/gin-shortner-api/service"
)

func main() {
	// initiate env
	Env := env.NewEnv()

	// initiate db connection
	db := app.NewDB(Env)

	// database migration
	err := db.AutoMigrate(&model.Shortener{})
	helper.PanicIfError(err)

	// initiate repository, service & controller
	shortenerRepository := repository.NewShortenerRepository(db)
	shortenerService := service.NewShortenerService(shortenerRepository)
	shortenerController := controller.NewShortenerController(shortenerService)

	// initiate router
	newRouter := router.NewRouter(shortenerController)

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// run and serve server
	err = newRouter.Run(":" + Env["APP_PORT"])
	helper.PanicIfError(err)
}
