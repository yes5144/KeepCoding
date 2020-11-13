package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOneVersionView(c *gin.Context) {
	c.HTML(http.StatusOK, "detail.html", nil)
}

func GetAllVersionView(c *gin.Context) {
	c.HTML(http.StatusOK, "allversions.html", nil)
}
