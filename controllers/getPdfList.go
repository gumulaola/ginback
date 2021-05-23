package controllers

import (
	"gin-app/initDB"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PdfList struct {
	PdfCourse string
	PdfChapter string
	PdfPath string
}

func GetPdfList(c *gin.Context)  {
	pdfLists := make([]PdfList, 0)
	rows, _ := initDB.DB.Query("select * from pdflist")
	for rows.Next() {
		var pdfList PdfList
		rows.Scan(&pdfList.PdfCourse, &pdfList.PdfChapter, &pdfList.PdfPath)
		pdfLists = append(pdfLists, pdfList)
	}
	c.JSON(http.StatusOK, pdfLists)
}
