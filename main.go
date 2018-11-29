package main

import (
	"./handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/solve", handlers.Solve)
	return router
}

func main() {
	r := SetupRouter()

	r.Run() // listen and serve on 0.0.0.0:8080
}
