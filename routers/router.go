package routers

import (
	"gin-app/controllers"
	"gin-app/controllers/education"
	"gin-app/controllers/miniprogram"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	// static files path
	r.StaticFS("/static", http.Dir("./upload"))

	// for mini-program
	r.GET("/", controllers.Index)
	r.GET("/getopenid", miniprogram.GetOpenIdHandle)
	r.GET("/getpostlist", miniprogram.GetUserPostList)
	r.GET("/create", miniprogram.Create)
	r.GET("/delete", miniprogram.Delete)

	// for education
	r.POST("/api/addcourse", education.AddCourse)
	r.GET("/api/getcourse", education.GetCourse)
	r.POST("/api/uploadfile", education.UploadFile)
	r.GET("/api/getpdflist", education.GetPdfList)
	return r
}
