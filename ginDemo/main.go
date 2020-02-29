package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yes5144/KeepCoding/ginDemo/controllers"
)

func main() {
	log.Println("hello gin")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	person := r.Group("/person")
	{
		person.GET("/all", controllers.QueryPersonAllView)
		person.POST("/new", controllers.InsertPersonView)
	}
	r.Run()
}
