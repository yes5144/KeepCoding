package controllers

import (
	"log"
	"net/http"

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
	ret, err := models.QueryPersonAll()
	if err != nil {
		log.Printf("query err, Failed code: %#v", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "查询成功",
		"data": ret,
	})
}
