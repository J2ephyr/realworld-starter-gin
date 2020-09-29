package common

import "github.com/gin-gonic/gin"

func Error(c *gin.Context, error string) {
	c.JSON(422, map[string]interface{}{
		"errors": error,
	})
}
