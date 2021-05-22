package routers

import (
	"gin-app/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/getopenid", controllers.GetOpenIdHandle)
	return r
}
