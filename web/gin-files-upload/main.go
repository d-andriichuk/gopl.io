package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		dst := fmt.Sprintf("/home/unknown/Documents/%s", file.Filename)

		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.String(500, fmt.Sprintf("%s file wasn't uploaded. Err: %v", file.Filename, err))
			fmt.Printf("%s file wasn't uploaded. Err: %v", file.Filename, err)
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("%s uploaded", file.Filename))
	})

	router.Run(":8080")
}
