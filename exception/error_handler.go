package exception

import (
	"errors"
	"fmt"
	"github.com/rendyuwu/gin-shortner-api/domain/web"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Recovery() func(c *gin.Context) {
	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		if notFoundError(c, err) {
			return
		}
		internalServerError(c, err)
	})
}

func ValidationError(c *gin.Context, err error) bool {
	if err != nil {
		var ver validator.ValidationErrors
		if errors.As(err, &ver) {
			c.JSON(http.StatusBadRequest, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   simple(ver),
			})
			return true
		}

		// unmarshal error
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
		})
		return true
	}
	return false
}

func notFoundError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}
		c.JSON(http.StatusNotFound, webResponse)
		return true
	} else {
		return false
	}
}

func simple(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs[f.Field()] = err
	}

	return errs
}

func internalServerError(c *gin.Context, err interface{}) {
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}
	c.JSON(http.StatusInternalServerError, webResponse)
}
