package education

import (
	"gin-app/initDB"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PdfList struct {
	PdfCourseID int
	PdfChapter  string
	PdfPath     string
}

func GetPdfList(c *gin.Context) {
	pdfLists := make([]PdfList, 0)
	rows, _ := initDB.DB.Query("select * from pdflist")
	for rows.Next() {
		var pdfList PdfList
		rows.Scan(&pdfList.PdfCourseID, &pdfList.PdfChapter, &pdfList.PdfPath)
		pdfLists = append(pdfLists, pdfList)
	}
	c.JSON(http.StatusOK, pdfLists)
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	courseID := c.PostForm("courseID")
	chapter := c.PostForm("chapter")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
	}
	c.SaveUploadedFile(file, "./upload/pdf/"+courseID+"-"+chapter+".pdf")

	// todo? connect the pdf list with the id of course
	initDB.DB.Exec("insert into pdflist (pdf_course_id, pdf_chapter, pdf_path) values (?, ?, ?)", courseID, chapter, "/upload/pdf/"+courseID+"-"+chapter+".pdf")
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}
