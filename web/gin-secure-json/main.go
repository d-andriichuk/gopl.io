package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lenf", "austin", "foo"}

		c.SecureJSON(http.StatusOK, names)
	})

	r.Run(":8080")
}
