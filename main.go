package main

import (
	"hospital-playlist/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// func corsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// 		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(200)
// 		} else {
// 			c.Next()
// 		}
// 	}
// }

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	// r.Use(corsMiddleware())

	routes.UserRoute(r)
	routes.UserDetailRoute(r)
	routes.UserProfileRoute(r)
	routes.DrugRoute(r)

	r.Run()
}
