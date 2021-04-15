package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)


func main() {
	// sh()
	// fmt.Println("Hello, World!")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		log.Println("hello")
		res := sh()
		c.JSON(200, gin.H{
			"message": res,
		})
	})
	router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			filename := filepath.Base(file.Filename)
			// Upload the file to specific dst.
			c.SaveUploadedFile(file, filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8010")
}
