package main

import (
	"./app"
)

func main() {
	r := app.SetupRouter()

	r.Run() // listen and serve on 0.0.0.0:8080
}
