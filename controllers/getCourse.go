package controllers

import (
	"gin-app/initDB"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CourseList struct {
	CourseID   int
	CourseName string
	CourseInfo string
}

func GetCourse(c *gin.Context) {
	courselists := make([]CourseList, 0)
	rows, _ := initDB.DB.Query("select * from courselist")
	for rows.Next() {
		var courseList CourseList
		rows.Scan(&courseList.CourseID, &courseList.CourseName, &courseList.CourseInfo)
		courselists = append(courselists, courseList)
	}
	c.JSON(http.StatusOK, courselists)
}
