package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "ok")
	fmt.Printf("%x\n", h.Sum(nil))

	// 直接使用md5 ew对象的Write方式也是一样的
	h2 := md5.New()
	h2.Write([]byte("The fog is getting thicker!"))
	h2.Write([]byte("ok"))
	fmt.Printf("%x\n", h2.Sum(nil))

	fmt.Println(GetMD5Hash("The fog is getting thicker!ok"))
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
