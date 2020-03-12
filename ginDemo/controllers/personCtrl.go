package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yes5144/KeepCoding/ginDemo/models"
)

func InsertPersonView(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	// rs, err := models.InsertPerson(firstName, lastName)
	err := models.InsertPerson(firstName, lastName)
	if err != nil {
		log.Printf("insert err, Failed code: %#v", err)
	}
	// log.Println(ret)
	c.JSON(http.StatusOK, gin.H{
		"msg": "成功insert",
	})
}

func QueryPersonAllView(c *gin.Context) {
	c.HTML(http.StatusOK, "person/userall.html", gin.H{
		"status": http.StatusOK,
		"title":  "userall",
		"msg":    "querey all success",
	})
	// c.HTML(http.StatusOK, "person/userall.html", nil)
}

func ApiQueryPersonAll(c *gin.Context) {
	ret, err := models.QueryPersonAll()
	if err != nil {
		log.Printf("quere all err, Failed code: %#v", err)
	}
	log.Println("query all: ", ret)
	c.JSON(http.StatusOK, gin.H{
		"data": ret,
	})
}

func QueryPersonOneView(c *gin.Context) {
	log.Println("query person one view")
	c.HTML(http.StatusOK, "person/userone.html", gin.H{
		"status": http.StatusOK,
		"title":  "queryOne",
		"msg":    "select one",
	})

}

func ApiQueryPersonOne(c *gin.Context) {
	log.Println("api query person one")
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Printf("api querey person one , Failed code: %#v", err)
	}
	ret, err := models.QueryPersonOne(id)
	if err != nil {
		log.Printf("query person One, Failed code: %#v", err)
	}
	log.Println("date ret: ", ret)
	c.JSON(http.StatusOK, gin.H{
		// "status": http.StatusOK,
		"data": ret,
	})
}
