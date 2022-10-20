package env

import (
	"github.com/joho/godotenv"
	"github.com/rendyuwu/gin-shortner-api/helper"
)

type MyEnv map[string]string

func NewEnv() MyEnv {
	myEnv, err := godotenv.Read()
	helper.PanicIfError(err)

	return myEnv
}
