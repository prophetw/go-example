package main

import (
	"fmt"
	"os/exec"
)
 
func img2pdf(filepaths []string, mergePdfName string) string {
	filepaths = append(filepaths, mergePdfName)	
	cmd := exec.Command("convert", filepaths...)
	err := cmd.Run()
	if err != nil {
			fmt.Println("Execute Command failed:" + err.Error())
			return ""
	}
	fmt.Println("Execute Command finished.")
	// 返回的应该是一条下载链接
	return mergePdfName
	// file, err := os.Open("1.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// buf := make([]byte, 126)
	// n, err := file.Read(buf)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("%d = %q", n, buf)
	// fmt.Println("Hello, sh!")
}

func mergePDF(filepaths []string, mergePdfName string) string {
	fmt.Println("start merge.")
	filepaths = append(filepaths, mergePdfName)	
	cmd := exec.Command("pdfunite", filepaths...)
	err := cmd.Run()
	if err != nil {
			fmt.Println("Execute Command failed:" + err.Error())
			return err.Error()
	}
	fmt.Println("Execute Command finished.")
	// 返回的应该是一条下载链接
	return mergePdfName
}