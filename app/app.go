package app

import (
	"./handlers/solve"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// Sets up router
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./app/public/home", true)))
	router.Use(static.Serve("/docs", static.LocalFile("./app/public/docs", true)))

	api := router.Group("/api")
	{
		api.POST("/solve", solve.Solve)
	}
	return router
}
