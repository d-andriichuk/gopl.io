package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/version", v1VersionEndpoint)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/version", v2VersionEndpoint)
	}

	router.Run(":8080")
}

func v1VersionEndpoint(c *gin.Context) {
	// c.String(http.StatusOK, "V1 api endpoint")
	c.JSON(http.StatusOK, gin.H{
		"content": "api version v1",
		"message": "processed",
	})
}

func v2VersionEndpoint(c *gin.Context) {
	// c.String(http.StatusOK, "V2 api endpoint")
	c.JSON(http.StatusOK, gin.H{
		"content": "api version v2",
		"message": "processed",
	})
}
