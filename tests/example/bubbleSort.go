package main

import "fmt"

func main() {
	arr2 := []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}

	BubbleSort(arr2)
	fmt.Println(arr2)
}
func BubbleSort(values []int) {
	for i := 1; i < len(values); i++ {
		for j := 0; j < len(values)-i; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
}
