package controllers

import (
	"gin-app/initDB"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	date := c.PostForm("date")
	openid := c.PostForm("openid")
	initDB.DB.Exec("insert into postlist (title, content, date, openid) values (?, ?, ?, ?)", title, content, date, openid)
}
