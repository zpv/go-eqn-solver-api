package app

import (
	"./handlers"
	"github.com/gin-gonic/gin"
)

// Sets up router
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/solve", handlers.Solve)
	return router
}
