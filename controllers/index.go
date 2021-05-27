package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   Data   `json:"data"`
}

type Data struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

func Index(c *gin.Context) {
	var data = Data{"Genre", "stu"}
	var response = Response{0, "success", data}
	c.JSON(http.StatusOK, response)
}
