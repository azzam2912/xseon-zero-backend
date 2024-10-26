package handler

import (
	"time"
	"xseon-zero/domain/response"

	"github.com/gin-gonic/gin"
)

func writeJSONResponse(c *gin.Context, msg string, data interface{}, statusCode int, timestamp time.Time) {
	response := response.Response{
		Msg:       msg,
		Data:      data,
		Timestamp: timestamp,
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(statusCode, response)
}
