package main

import (
	"fmt"
	"strings"

	"common.dh.cn/utils"
)

func main() {
	str := `[["(to_char(CURRENT_TIMESTAMP, 'YYYY-MM-DD HH24:MI:SS'))","(to_char(CURRENT_TIMESTAMP, 'YYYY-MM-DD HH24:MI:SS'))"]]`
	fmt.Println(str)
	values := make([]string, 0)
	data := strings.Split(strings.Trim(strings.Trim(str, "["), "]"), `","`)
	for _, val := range data {
		val = strings.Trim(val, `"`)
		if utils.IsEmpty(val) {
			continue
		}

		values = append(values, val)
	}
	fmt.Println(values)

	var a float64 = 7
	var b float64 = 3
	fmt.Println(a / b)
}
