package main

import (
	"fmt"
)

func quickSort(values []int, left int, right int) {
	if left < right {

		// 设置基准值
		temp := values[left]

		// 设置哨兵
		i, j := left, right
		for {
			// 从右向左找，找到第一个比基准值小的数
			for values[j] >= temp && i < j {
				j--
			}

			// 从左向右找，找到第一个比基准值大的数
			for values[i] <= temp && i < j {
				i++
			}

			// 如果哨兵相遇，则退出循环
			if i >= j {
				break
			}

			// 交换左右两侧的值
			values[i], values[j] = values[j], values[i]
		}

		// 将基准值移到哨兵相遇点
		values[left] = values[i]
		values[i] = temp

		// 递归，左右两侧分别排序
		quickSort(values, left, i-1)
		quickSort(values, i+1, right)
	}
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}

func BubbleSort(values []int) {
	flag := true
	for i := 0; i < len(values)-1; i++ {
		flag = true
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}
		if flag == true {
			// 如果已经顺序对了，就不用继续冒泡排序了。
			break
		}
	}
}

func main() {
	arr := []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	QuickSort(arr)
	fmt.Println(arr)

	arr2 := []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	BubbleSort(arr2)
	fmt.Println(arr2)
}
