package handlers

import "github.com/gin-gonic/gin"

func Solve(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
