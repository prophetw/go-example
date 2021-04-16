package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
			if c.Request.Method == "OPTIONS" {
					c.AbortWithStatus(204)
					return
			}
			c.Next()
	}
}

func main() {
	// sh()
	fmt.Println("Hello, World!")
	// userType := &user.User{"0", "0", "ubuntu", "",""}
	pagesize,_ := user.Current()
	// hostname := os.Hostname()
	fmt.Println(*pagesize)
	// fmt.Println(hostname)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// serve static files
	router.Static("/uploads", "./uploads") 

	router.Use(CORSMiddleware())
	router.GET("/ping", func(c *gin.Context) {
		log.Println("hello")
		res := sh()
		c.JSON(200, gin.H{
			"message": res,
		})
	})
	router.GET("/", func(c *gin.Context) {
		log.Println("hello")
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		path, _ := os.Getwd()
		log.Println("upload")
		form, _ := c.MultipartForm()
		files := form.File["file[]"]
		for _, file := range files {
			log.Println(file)
			filename := filepath.Base(file.Filename)
			// Upload the file to specific dst.
			dstFolderPath :=path+"/uploads" 
			if _, err := os.Stat(dstFolderPath); os.IsNotExist(err) {
				os.Mkdir(dstFolderPath, 0755)
			}
			log.Println(dstFolderPath)
			err := c.SaveUploadedFile(file, dstFolderPath + "/" + filename)
			if err != nil {
				log.Println(err)
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8010")
}
