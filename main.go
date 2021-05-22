package main

import (
	"gin-app/routers"
)

func main() {
	r := routers.InitRouter()
	r.Run()

	// r := gin.Default()

	// index
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello",
	// 	})
	// })

	// upload
	// r.POST("/upload", func(c *gin.Context) {
	// 	file, _ := c.FormFile("file")
	// 	log.Println(file.Filename)
	// 	filePath := "upload/"
	// 	path := filePath + file.Filename
	// 	c.SaveUploadedFile(file, path)
	// 	c.String(http.StatusOK, fmt.Sprintf("uploaded!"))
	// })

	// r.Run()
}
