package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"

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
	// res := sh()
	fmt.Println("Hello, World!")
	dstFolderPath := "/var/www/pdfutils/uploads"
	// userType := &user.User{"0", "0", "ubuntu", "",""}
	pagesize, _ := user.Current()
	// hostname := os.Hostname()
	fmt.Println(*pagesize)
	// fmt.Println(hostname)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// serve static files
	router.Static("/uploads", "/var/www/pdfutils/uploads")
	router.Use(CORSMiddleware())
	router.GET("/ping", func(c *gin.Context) {
		log.Println("hello")
		// res := sh()
		c.JSON(200, gin.H{
			// "message": res,
			"message": "res",
		})
	})
	router.GET("/", func(c *gin.Context) {
		log.Println("hello")
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/imgtopdf", func(c *gin.Context) {
		// Multipart form
		log.Println("upload")
		form, _ := c.MultipartForm()
		files := form.File["file[]"]
		filepaths := []string{}
		mergeName := ""
		for _, file := range files {
			filename := filepath.Base(file.Filename)
			// Upload the file to specific dst.
			if _, err := os.Stat(dstFolderPath); os.IsNotExist(err) {
				os.Mkdir(dstFolderPath, 0755)
			}
			log.Println(dstFolderPath)
			err := c.SaveUploadedFile(file, dstFolderPath+"/"+filename)
			if err != nil {
				log.Println(err)
			}
			filepaths = append(filepaths, dstFolderPath+"/"+filename)
			filenameWithSuffix := path.Base(filename)
			fileSuffix := path.Ext(filenameWithSuffix)
			filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
			mergeName += filenameOnly
		}
		// merge
		res := img2pdf(filepaths, dstFolderPath+"/"+mergeName+".pdf")
		log.Println(" ---- hhhh ---- ")
		log.Println(filepaths)
		log.Println(mergeName)
		if res == "" {
			c.JSON(200, gin.H{
				"code":    900100,
				"message": "转换失败",
			})
		} else {
			c.JSON(200, gin.H{
				"code":    200,
				"name":    mergeName + ".pdf",
				"message": res,
				"link":    "/uploads/" + mergeName + ".pdf",
			})
		}

	})
	router.POST("/pdfmerge", func(c *gin.Context) {
		// Multipart form
		log.Println("upload")
		form, _ := c.MultipartForm()
		files := form.File["file[]"]
		filepaths := []string{}
		mergeName := "merge"
		for _, file := range files {
			filename := filepath.Base(file.Filename)
			// Upload the file to specific dst.
			if _, err := os.Stat(dstFolderPath); os.IsNotExist(err) {
				os.Mkdir(dstFolderPath, 0755)
			}
			log.Println(dstFolderPath)
			err := c.SaveUploadedFile(file, dstFolderPath+"/"+filename)
			if err != nil {
				log.Println(err)
			}
			filepaths = append(filepaths, dstFolderPath+"/"+filename)
			filenameWithSuffix := path.Base(filename)
			fileSuffix := path.Ext(filenameWithSuffix)
			filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
			mergeName += filenameOnly
		}
		// merge
		res := mergePDF(filepaths, dstFolderPath+"/"+mergeName+".pdf")
		log.Println(" ---- hhhh ---- ")
		log.Println(filepaths)
		log.Println(mergeName)
		if res == "" {
			c.JSON(200, gin.H{
				"code":    900100,
				"message": "转换失败",
			})
		} else {
			c.JSON(200, gin.H{
				"code":    200,
				"name":    mergeName + ".pdf",
				"message": res,
				"link":    "/uploads/" + mergeName + ".pdf",
			})
		}
	})
	router.Run(":8010")
}
