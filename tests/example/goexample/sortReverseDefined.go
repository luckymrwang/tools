package main

import (
	"fmt"
	"sort"
)

// 自定义的 Reverse 类型
type Reverse struct {
	sort.Interface // 这样， Reverse 可以接纳任何实现了 sort.Interface (包括 Len, Less, Swap 三个方法) 的对象
}

// Reverse 只是将其中的 Inferface.Less 的顺序对调了一下
func (r Reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

func main() {
	ints := []int{5, 2, 6, 3, 1, 4} // 未排序

	sort.Ints(ints)                            // 特殊排序函数， 升序
	fmt.Println("after sort by Ints:\t", ints) // [1 2 3 4 5 6]

	doubles := []float64{2.3, 3.2, 6.7, 10.9, 5.4, 1.8}

	sort.Float64s(doubles)                            // float64 排序版本 1
	fmt.Println("after sort by Float64s:\t", doubles) // [1.8 2.3 3.2 5.4 6.7 10.9]

	strings := []string{"hello", "good", "students", "morning", "people", "world"}
	sort.Strings(strings)
	fmt.Println("after sort by Strings:\t", strings) // [good hello mornig people students world]

	ipos := sort.SearchInts(ints, -1)       // int 搜索
	fmt.Printf("pos of 5 is %d th\n", ipos) // 并不总是正确呀 ! (搜索不是重点)

	dpos := sort.SearchFloat64s(doubles, 20.1) // float64 搜索
	fmt.Printf("pos of 5.0 is %d th\n", dpos)  // 并不总是正确呀 !

	fmt.Printf("doubles is asc ? %v\n", sort.Float64sAreSorted(doubles))

	doubles = []float64{3.5, 4.2, 8.9, 100.98, 20.14, 79.32}
	// sort.Sort(sort.Float64Slice(doubles))    // float64 排序方法 2
	// fmt.Println("after sort by Sort:\t", doubles)    // [3.5 4.2 8.9 20.14 79.32 100.98]
	(sort.Float64Slice(doubles)).Sort()           // float64 排序方法 3
	fmt.Println("after sort by Sort:\t", doubles) // [3.5 4.2 8.9 20.14 79.32 100.98]

	sort.Sort(Reverse{sort.Float64Slice(doubles)})         // float64 逆序排序
	fmt.Println("after sort by Reversed Sort:\t", doubles) // [100.98 79.32 20.14 8.9 4.2 3.5]
}
