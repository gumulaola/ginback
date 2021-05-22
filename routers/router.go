package routers

import (
	"gin-app/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", controllers.Index)
	r.GET("/getopenid", controllers.GetOpenIdHandle)
	r.GET("/getpostlist", controllers.GetUserPostList)
	r.POST("/create", controllers.Create)
	return r
}
