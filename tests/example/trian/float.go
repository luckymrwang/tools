package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fmt.Sprintf("%s", fmt.Sprintf("%f", float64(13151332654))))
	fmt.Println(format12(float64(13151332658.4)))
	fmt.Println(float64(131513326584.123445))
	fmt.Println(strconv.FormatFloat(float64(131513326584.123456), 'f', -1, 64))
}

func format12(x float64) string {
	if x >= 1e12 {
		// Check to see how many fraction digits fit in:
		s := fmt.Sprintf("%.g", x)
		format := fmt.Sprintf("%%12.%dg", 12-len(s))
		return fmt.Sprintf(format, x)
	}

	// Check to see how many fraction digits fit in:
	s := fmt.Sprintf("%.0f", x)
	if len(s) == 12 {
		return s
	}
	format := fmt.Sprintf("%%%d.%df", len(s), 12-len(s)-1)
	return fmt.Sprintf(format, x)
}
