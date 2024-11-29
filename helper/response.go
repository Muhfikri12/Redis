package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPResponse struct {
	Status    bool        `json:"status"`
	ErrorCode string      `json:"error_msg,omitempty"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

func ResponseOK(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, HTTPResponse{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func ResponseError(c *gin.Context, errorCode string, message string, httpStatusCode int) {
	c.JSON(httpStatusCode, HTTPResponse{
		Status:    false,
		ErrorCode: errorCode,
		Message:   message,
	})
}
