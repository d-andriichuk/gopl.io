package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123

		c.JSON(http.StatusOK, msg)
	})

	router.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	router.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"content": gin.H{"user": "Toha", "message": "hey"}, "status": http.StatusOK})
	})

	router.Run(":8080")
}
