package miniprogram

import (
	"gin-app/initDB"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Delete(c *gin.Context) {
	postid := c.Query("postid")
	openid := c.Query("openid")
	initDB.DB.Exec("DELETE FROM postlist WHERE postid = ? AND openid = ?", postid, openid)
	c.JSON(http.StatusOK, gin.H{
		"code": "0",
		"msg":  "OK",
	})
}
