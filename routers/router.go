package routers

import (
	"gin-app/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controllers.Index)
	r.GET("/getopenid", controllers.GetOpenIdHandle)
	r.GET("/getpostlist", controllers.GetUserPostList)
	r.POST("/create", controllers.Create)
	return r
}
