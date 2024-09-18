package main

import (
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/img/profile", "./img/profile")
	r.Use(corsMiddleware())
	routers.RouterCombine(r)
	r.Run("0.0.0.0:8080")
}
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, POST, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
