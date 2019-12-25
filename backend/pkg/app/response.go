package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/gin_todo_demo/backend/pkg/e"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse send custom response
func SendResponse(c *gin.Context, httpCode int, err error, data interface{}) {
	code, message := e.DecodeErr(err)
	
	c.JSON(httpCode, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
