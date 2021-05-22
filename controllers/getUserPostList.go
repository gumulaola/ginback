package controllers

import (
	"gin-app/initDB"
	"github.com/gin-gonic/gin"
	"net/http"
)

type List struct {
	Title string
	Content string
	Date string
	Openid string
}

func GetUserPostList(c *gin.Context) {
	code := c.Query("code")
	lists := make([]List, 0)
	rows, _ := initDB.DB.Query("select * from postlist where openid = ?", code)
	for rows.Next() {
		var list List
		rows.Scan(&list.Title, &list.Content, &list.Date, &list.Openid)
		lists = append(lists, list)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": lists,
	})
}
