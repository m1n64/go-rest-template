package actions

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Ping(c *gin.Context) {

	c.JSON(200, gin.H{
		"success":   true,
		"timestamp": time.Now(),
		"message":   "pong",
	})
}
