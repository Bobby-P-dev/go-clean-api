package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func Error(c *gin.Context, statusCode int, message string, err []string) {
	c.JSON(statusCode, gin.H{
		"success": false,
		"message": message,
		"error":   err,
	})
}

func SuccsesMeta(c *gin.Context, message string, data interface{}, meta interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message,
		"data":    data,
		"meta":    meta,
	})
}

func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}
