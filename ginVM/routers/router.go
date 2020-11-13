package routers

import (
	"github.com/gin-gonic/gin"
	// v1 "github.com/yes5144/KeepCoding/ginVM/controllers/api/v1"

	"ginVM/controllers"
	v1 "ginVM/controllers/api/v1"
)

func InitRouters() *gin.Engine {
	// r := gin.Default()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode("debug")
	r.LoadHTMLGlob("templates/**/*")
	// version
	verView := r.Group("/version")
	{
		verView.GET("/detail/:id", controllers.GetOneVersionView)
		verView.GET("/", controllers.GetAllVersionView)
	}

	apiv1 := r.Group("/api/v1")
	{
		verapi := apiv1.Group("/version")
		{
			verapi.GET("/all", v1.GetAllVersionApi)
			verapi.GET("/get/:id", v1.GetOneVersionApi)
			verapi.GET("/new", v1.NewOneVersionApi)
		}
	}
	return r
}
