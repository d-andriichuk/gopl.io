package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gopl.io/grpc/metrics/service"
)

const (
	addr = "localhost:50051"
)

var (
	metrics *service.Metrics
)

func main() {
	metrics = service.NewMetrics()

	listen()
}

func listen() {
	router := gin.Default()

	router.POST("/metrics", metrixHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
}

func metrixHandler(c *gin.Context) {
	if err := metrics.Exchange(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed to exchange metrcis with persister",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "metrics saved",
		"error":  "nil",
	})
}
