package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	// r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
	r.GET("/benchmark", benchEndpoint)

	authorized := r.Group("/")

	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)

		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	r.Run(":8080")
}

func benchEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content": "bench",
		"message": "processed",
	})
}
