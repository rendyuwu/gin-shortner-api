package helper

import (
	"errors"
	"github.com/rendyuwu/gin-shortner-api/exception"
	"gorm.io/gorm"
)

func IsNotFoundError(err error) {
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(exception.NewNotFoundError("book is not found"))
		} else {
			panic(err)
		}
	}
}
