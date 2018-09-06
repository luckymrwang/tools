package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xlsx, err := excelize.OpenFile("/Users/sino/Downloads/tpc/a.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("工作表1")
	for i, row := range rows {
		xlsx.SetCellValue("工作表1", fmt., "Hello world.")
		fmt.Println()
	}
}
