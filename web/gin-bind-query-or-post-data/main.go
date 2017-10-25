package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Person struc
type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	router := gin.Default()
	router.POST("/testing", startPage)
	router.Run(":8080")
}

func startPage(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.String(http.StatusBadRequest, "error: %v", err.Error())
		return
	}

	log.Println(person.Name)
	log.Println(person.Address)
	log.Println(person.Birthday)

	c.String(http.StatusOK, "Success")
}
