package main

import "github.com/gin-gonic/gin"
import "net/http"

// LoginForm struct
type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var form LoginForm
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if form.User == "user" && form.Password == "password" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unathorized"})
	})
	router.Run(":8080")
}
