package controllers

import (
	"gin-app/initDB"
	"github.com/gin-gonic/gin"
)

func AddCourse(c *gin.Context) {
	courseName := c.PostForm("coursename")
	courseInfo := c.PostForm("courseinfo")
	initDB.DB.Exec("insert into courselist (coursename, courseinfo) values (?, ?)", courseName, courseInfo)
}
