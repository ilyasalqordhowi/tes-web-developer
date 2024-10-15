package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerOK(c *gin.Context, msg string, data any, pageInfo any) {
	c.JSON(http.StatusOK, Message{
		Success:  true,
		Message:  msg,
		Results:   data,
	})
}


func HandlerNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Message{
		Success: false,
		Message: message,
	})
}

func HandlerBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Message{
		Success: false,
		Message: message,
	})
}

func HandlerMaxFile(c *gin.Context, message string) {
	c.JSON(http.StatusRequestEntityTooLarge, Message{
		Success: false,
		Message: message,
	})
}