package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	// float4List := [] float32 {4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}    // no function : sort.Float32s
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}
	stringList2 := []string{"1", "3", "2", "12", "6", "9", "21", "32", "101", "102"}

	sort.Ints(intList)
	sort.Float64s(float8List)
	sort.Strings(stringList)
	sort.Strings(stringList2)

	fmt.Printf("%v\n%v\n%v\n%v\n", intList, float8List, stringList, stringList2)

	slic, err := stringToInt(stringList2)
	if err != nil {
		return
	}

	//	sort.Ints(slic)
	sort.Sort(sort.Reverse(sort.IntSlice(slic)))
	fmt.Printf("%v\n", slic)
}

func stringToInt(s []string) ([]int, error) {
	arr := make([]int, len(s))
	for k, v := range s {
		val, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}

		arr[k] = int(val)
	}

	return arr, nil
}
