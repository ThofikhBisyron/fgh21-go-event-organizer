package main

import (
	"fmt"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	password := "12345678"
	hashedPassword := lib.Encrypt(password)
	fmt.Println("Hashed password:", hashedPassword)
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
