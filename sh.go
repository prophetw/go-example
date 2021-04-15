package main

import (
	"fmt"
	"os/exec"
)
func sh() string {

	fileA := "222.jpg"
	fileB := "223.jpg"
	mergePdfName := fileA + fileB + ".pdf"
	params := []string{fileA, fileB, mergePdfName}

	cmd := exec.Command("convert", params...)
	err := cmd.Run()
	if err != nil {
			fmt.Println("Execute Command failed:" + err.Error())
			return ""
	}
	fmt.Println("Execute Command finished.")
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
