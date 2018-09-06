package main

import (
	"fmt"

	"common.dh.cn/connecters"
)

func main() {
	fmt.Println("begin csv connecter testing......")
	xlsx := connecters.NewExcel("/Users/sino/Downloads/aaa.xlsx")
	err := xlsx.ReadAll(true, true)
	if err != nil {
		fmt.Println(err)
	}
	a := xlsx.Data["客户信息表"]
	fmt.Println(a[0], len(a[0]))
}
