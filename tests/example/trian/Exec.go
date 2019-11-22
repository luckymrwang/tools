package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("/Users/sino/Documents/luckymrwang/go/src/xiaomi.dh.cn/xiaomi.dh.cn", "-service=Test2", "-method=FuncFour")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
	str := fmt.Sprintf("%v:%v", stdout.String(), stderr.String())
	fmt.Println(str)
}
