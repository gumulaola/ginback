package controllers

import (
	"gin-app/initDB"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	title := c.Query("title")
	content := c.Query("content")
	date := c.Query("date")
	openid := c.Query("openid")
	initDB.DB.Exec("insert into postlist (title, content, date, openid) values (?, ?, ?, ?)", title, content, date, openid)
}
