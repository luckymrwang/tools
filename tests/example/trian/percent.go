package main

import (
	"fmt"
)

func main() {
	a := GetCompareRate(float64(290), float64(878))
	fmt.Println(a)
}

func GetCompareRate(newData, oldData float64) string {
	if newData == 0 && oldData == 0 {
		return "0%"
	} else if newData != 0 && oldData == 0 {
		return "100%"
	}

	rateNum := fmt.Sprintf("%.1f", ((newData/oldData)-1)*100)
	return rateNum + "%"
}
