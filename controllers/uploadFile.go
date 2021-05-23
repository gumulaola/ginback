package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	course := c.PostForm("course")
	chapter := c.PostForm("chapter")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
	}
	c.SaveUploadedFile(file, "./upload/pdf/"+course+"-"+chapter+".pdf")
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}
