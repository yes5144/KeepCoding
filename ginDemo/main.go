package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yes5144/KeepCoding/ginDemo/controllers"
)

func main() {
	log.Println("hello gin")
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/ping", func(c *gin.Context) {
		// c.JSON(http.StatusOK, gin.H{"msg": "pong"})
		c.HTML(http.StatusOK, "upload/index.html", nil)
	})

	// upload
	r.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload"]
		today := time.Now().Format("20060102")
		log.Println(today)
		for _, file := range files {
			log.Println(file.Filename)
			// 上传到指定目录
			uploadDir := "upload/" + today + "/"
			info, err := os.Stat(uploadDir)
			log.Println("info of dir:", info, err)
			if err != nil {
				log.Printf("目标文件夹不存在, Failed code: %#v", err)
				if os.IsNotExist(err) {
					os.Mkdir(uploadDir, os.ModePerm)
				}
			}
			dst := uploadDir + file.Filename
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"上传文件数量：": len(files),
		})
	})
	person := r.Group("/person")
	{
		person.GET("/all", controllers.QueryPersonAllView)
		person.GET("/one", controllers.QueryPersonOneView)
		person.POST("/new", controllers.InsertPersonView)
		// person.POST("/user", controllers.InsertPersonView)
	}
	api := r.Group("/v1")
	{
		api.GET("/person/one", controllers.ApiQueryPersonOne)
		api.GET("/person/all", controllers.ApiQueryPersonAll)
	}
	r.Run()
}

// https://gin-gonic.com/zh-cn/docs/
