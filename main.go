package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Id       int
	Name     string
	Email    string
	Password string
}
type Server struct {
	Succes  bool
	Message string
	Results []Data
}

func main() {
	r := gin.Default()
	r.Use(corsMiddleware())
	data := []Data{{
		Id:       1,
		Name:     "Fazztrack",
		Email:    "Rais@yahoo.com",
		Password: "1234",
	},
		{
			Id:       2,
			Name:     "admin",
			Email:    "admin@mail.com",
			Password: "1234",
		},
	}

	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, []Server{
			Server{
				Succes:  true,
				Message: "Ok",
				Results: data,
			},
		})
	})

	r.POST("/users", func(c *gin.Context) {
		user := Data{}

		c.Bind(&user)
		user.Id = len(data) + 1

		data = append(data, user)

		c.JSON(http.StatusOK, []Server{
			Server{
				Succes:  true,
				Message: "New Data",
				Results: data,
			},
		})
	})
	r.POST("/auth/login", func(c *gin.Context) {
		userlogin := Data{}

		c.Bind(&userlogin)
		email := userlogin.Email
		password := userlogin.Password

		dataResults := true
		if dataResults {
			for dataResults {
				for i := 0; i < len(data); i++ {
					resultsEmail := data[i].Email
					resultsPassword := data[i].Password
					if email == resultsEmail && password == resultsPassword {
						c.JSON(http.StatusOK, Server{
							Succes:  true,
							Message: "Login succes",
						})
						return
					}
				}

				dataResults = false
			}

			c.JSON(http.StatusUnauthorized, Server{
				Succes:  false,
				Message: "Wrong email or password",
			})
		}
	})
	r.PATCH("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, Server{
				Succes:  false,
				Message: "Wrong Text In Url/Endpoint",
			})
			return
		}

		updatedUser := Data{}
		if err := c.Bind(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, Server{
				Succes:  false,
				Message: "Server Down!!",
			})
			return
		}

		for i, u := range data {
			if u.Id == id {
				data[i].Name = updatedUser.Name
				data[i].Email = updatedUser.Email
				data[i].Password = updatedUser.Password
				c.JSON(http.StatusOK, Server{
					Succes:  true,
					Message: "User Updated",
					Results: data,
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, Server{
			Succes:  false,
			Message: "User Not Found",
		})
	})
	r.DELETE("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, Server{
				Succes:  false,
				Message: "ID Not Found",
			})
			return
		}

		for i, u := range data {
			if u.Id == id {
				data = append(data[:i], data[i+1:]...)
				c.JSON(http.StatusOK, Server{
					Succes:  true,
					Message: "User deleted successfully",
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, Server{
			Succes:  false,
			Message: "User not found",
		})
	})
	// r.GET("/users/:id", func(c *gin.Context) {
	// 	idStr := c.Param("id")
	// 	id, _ := strconv.Atoi(idStr)

	// 	for getId  := range data {
	// 		if getId.Id == id {
	// 			c.JSON(http.StatusOK, Server{
	// 				Succes:  true,
	// 				Message: "nama ketemu",
	// 				Results: []Data{getId},
	// 			})
	// 			return
	// 		}
	// 	}

	// 	c.JSON(http.StatusNotFound, Server{
	// 		Succes:  false,
	// 		Message: "Name Not Found",
	// 	})
	// })

	r.Run("localhost:8080")
}
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
