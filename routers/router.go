package routers

import (
	"gin-app/controllers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.StaticFS("/static", http.Dir("./upload"))
	r.GET("/", controllers.Index)
	r.GET("/getopenid", controllers.GetOpenIdHandle)
	r.GET("/getpostlist", controllers.GetUserPostList)
	r.GET("/create", controllers.Create)
	r.POST("/addcourse", controllers.AddCourse)
	r.GET("/getcourse", controllers.GetCourse)
	r.POST("/uploadfile", controllers.UploadFile)
	return r
}
