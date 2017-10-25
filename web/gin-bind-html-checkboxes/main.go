package main

import (
	"github.com/gin-gonic/gin"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("views/*")
	router.GET("/", indexHandler)
	router.POST("/", formHandler)

	router.Run(":8080")
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.Bind(&fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}
