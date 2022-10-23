//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/rendyuwu/gin-shortner-api/app"
	"github.com/rendyuwu/gin-shortner-api/controller"
	"github.com/rendyuwu/gin-shortner-api/env"
	"github.com/rendyuwu/gin-shortner-api/repository"
	"github.com/rendyuwu/gin-shortner-api/router"
	"github.com/rendyuwu/gin-shortner-api/service"
)

var shortenerSet = wire.NewSet(repository.NewShortenerRepository, service.NewShortenerService, controller.NewShortenerController)

func InitializedServer() *gin.Engine {
	wire.Build(env.NewEnv, app.NewDB, router.NewRouter, shortenerSet)

	return nil
}
