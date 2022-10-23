package app

import (
	"github.com/rendyuwu/gin-shortner-api/domain/model"
	"github.com/rendyuwu/gin-shortner-api/env"
	"github.com/rendyuwu/gin-shortner-api/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(env env.MyEnv) *gorm.DB {
	dsn := "" + env["MYSQL_USER"] + "" + env["MYSQL_PASSWORD"] + ":@tcp(" + env["MYSQL_HOST"] + ":" + env["MYSQL_PORT"] + ")/" + env["MYSQL_DATABASE"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	// database migration
	err = db.AutoMigrate(&model.Shortener{})
	helper.PanicIfError(err)

	return db
}
