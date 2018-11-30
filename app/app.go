package app

import (
	"./handlers/solve"
	"github.com/gin-gonic/gin"
)

// Sets up router
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/solve", solve.Solve)
	return router
}
