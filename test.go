package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	
	publicIp := "/home/ubuntu/publicIp.txt"

	file, err := os.Open(publicIp)
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
			log.Fatal(err)
	}
	lastRecordIpStr := string(data)
	fmt.Println("lastIp: \n", lastRecordIpStr)
	resp, _ := http.Get("http://ip.cip.cc/")
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	curIpStr := string(body)
	fmt.Println("curIpStr: \n", curIpStr)
	if curIpStr == lastRecordIpStr{
		fmt.Println("same")
	}else{
		file, err := os.OpenFile(
			publicIp,
			os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
			0666,
		)
		if err != nil {
				log.Fatal(err)
		}
		defer file.Close()
		byteSlice := []byte(curIpStr)
		_, err = file.Write(byteSlice)
    if err != nil {
        log.Fatal(err)
    }
		curIp := strings.TrimSuffix(curIpStr, "\n")
		// newPara := []string{"jjjjj", ">>", publicIp}
		// cmd := exec.Command("echo", params...)
		cmds := "echo " + curIp + "| mutt -s ipChanged 532300391@qq.com" 
		cmd := exec.Command("bash", "-c", cmds)
		cmd.Run()
	}
}
