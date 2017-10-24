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
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Printf("File to upload: %s\n", file.Filename)

			dst := fmt.Sprintf("/home/unknown/Documents/%s", file.Filename)

			if err := c.SaveUploadedFile(file, dst); err != nil {
				log.Printf("File %s wasn't upladed. Err: %v\n", file.Filename, err)
				c.String(500, "File %s wasn't uploaded. Err: %v\n", file.Filename, err)
				return
			}

			c.String(http.StatusOK, "File %s was uploaded\n", file.Filename)
		}
	})

	router.Run(":8080")
}
