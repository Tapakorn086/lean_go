package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UserResponse struct {
	FirstName string `json:"frist_name"`
	LastName  string `json:"last_name"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "welcome to GIN"})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World"})
	})

	r.GET("/user/:username", func(c *gin.Context) {
		username := c.Param("username")
		c.JSON(http.StatusOK, gin.H{"username": username})
	})

	r.POST("/users", func(ctx *gin.Context) {
		var user User
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response := UserResponse{
			FirstName: user.Name,
			LastName:  "Nakason",
		}
		ctx.JSON(http.StatusOK, response)
	})

	adminGroup := r.Group("/admin")
	{
		adminGroup.GET("/dashboard", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome to admin dashboard page"})
		})
	}

	r.Run()
}
