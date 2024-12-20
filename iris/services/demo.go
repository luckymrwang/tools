package services

import (
	"fmt"

	"tools/iris/common"

	"github.com/kataras/iris/v12"
)

type DemoService struct {
	Ctx iris.Context
}

func GetDemoService(ctx iris.Context) *DemoService {
	return &DemoService{Ctx: ctx}
}

type A struct {
	Name string
	Age  int
}

func (s *DemoService) Echo() error {
	a := &A{}
	if common.IsNil(a) {
		fmt.Println("the first is null")
	} else {
		fmt.Println("the first is not null")
	}
	var b A
	if common.IsNil(b) {
		fmt.Println("the second is null")
	} else {
		fmt.Println("the second is not null")
	}
	if common.IsNil(nil) {
		fmt.Println("the third is null")
	} else {
		fmt.Println("the third is not null")
	}

	return nil
}

func difference(a, b []int) []int {
	m := make(map[int]struct{}, len(b))
	for _, v := range b {
		m[v] = struct{}{}
	}
	var diff []int
	for _, v := range a {
		if _, found := m[v]; !found {
			diff = append(diff, v)
		}
	}
	return diff
}

func SymmetricDifference(slice1, slice2 []int) []int {
	diff1 := difference(slice1, slice2)
	diff2 := difference(slice2, slice1)
	return append(diff1, diff2...)
}
