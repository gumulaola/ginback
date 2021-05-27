package education

import (
	"gin-app/initDB"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CourseList struct {
	CourseID   int    `json:"course_id"`
	CourseName string `json:"course_name"`
	CourseInfo string `json:"course_info"`
}

func AddCourse(c *gin.Context) {
	courseName := c.PostForm("coursename")
	courseInfo := c.PostForm("courseinfo")
	initDB.DB.Exec("insert into courselist (coursename, courseinfo) values (?, ?)", courseName, courseInfo)
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
