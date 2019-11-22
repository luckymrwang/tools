package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"time"

	"common.dh.cn/constants"
)

const maxSize = 5 << 20 // 5MB

func main() {
	fmt.Println(maxSize)
	cmd := exec.Command("pwd", "-v")
	cmd.Dir = "/Users/sino/Documents/"
	_, err := doExec(cmd, "--help")
	fmt.Println(err)
}

func doExec(r *exec.Cmd, cmd string, exp ...int) (str string, e error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	r.Stdout = &stdout
	r.Stderr = &stderr
	fmt.Println("Exec Start:", cmd)
	r.Start()
	if len(exp) < 1 {
		exp = []int{constants.IMPORT_EXPIRE}
	}
	done := make(chan error)
	go func() { done <- r.Wait() }()

	timeout := time.After(time.Duration(exp[0]) * time.Second)
	select {
	case <-timeout:
		r.Process.Kill()
		e = errors.New("Command timed out")
	case e = <-done:
		str = fmt.Sprintf("%v:%v", stdout.String(), stderr.String())
	}
	if e != nil {
		errStr := fmt.Sprintf("Exec Failed:%s %s %s", e.Error(), str, cmd)
		fmt.Println(errStr)
		e = fmt.Errorf(errStr)
	} else {
		fmt.Println(fmt.Sprintf("Exec End:[%v]", stdout.String()), cmd)
	}

	return
}
