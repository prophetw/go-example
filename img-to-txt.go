package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func imgToTxt(imgPath string) string{
	cmds := "tesseract " + imgPath +" stdout --oem 1 -l chi_sim"
	cmd := exec.Command("bash", "-c", cmds)
	stdout, _ := cmd.StdoutPipe()   //创建输出管道
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		log.Fatalf("cmd.Start: %v")
	}
	fmt.Println("hellll ")
	result, _ := ioutil.ReadAll(stdout) // 读取输出结果
	resdata := string(result)
	fmt.Println(resdata)
	return resdata
}
