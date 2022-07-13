package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorRsp struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Level   int    `json:"level"`
}

func ErrorResponse(err error) ErrorRsp {
	return ErrorRsp{
		Message: err.Error(),
		Code:    0,
	}
}

func responseObject(c *gin.Context, out any, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
	} else {
		if out == nil {
			c.JSON(http.StatusNotFound, struct{}{})
		} else {
			c.JSON(http.StatusOK, out)
		}
	}
}
