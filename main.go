package main

import (
	"github.com/rendyuwu/gin-shortner-api/env"
	"github.com/rendyuwu/gin-shortner-api/helper"
)

func main() {
	// initiate env
	Env := env.NewEnv()

	// initiate router
	newRouter := InitializedServer()

	// run and serve server
	err := newRouter.Run(":" + Env["APP_PORT"])
	helper.PanicIfError(err)
}
