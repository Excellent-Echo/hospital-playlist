package main

import (
	"hospital-playlist/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.UserRoute(r)

	// r.Run()
	r.Run(":4444")
}
