package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string // 姓名
	Age  int    // 年纪
}

type PersonWrapper struct {
	people []Person
	by     func(p, q *Person) bool
}

func (pw PersonWrapper) Len() int { // 重写 Len() 方法
	return len(pw.people)
}
func (pw PersonWrapper) Swap(i, j int) { // 重写 Swap() 方法
	pw.people[i], pw.people[j] = pw.people[j], pw.people[i]
}
func (pw PersonWrapper) Less(i, j int) bool { // 重写 Less() 方法
	return pw.by(&pw.people[i], &pw.people[j])
}

func main() {
	people := []Person{
		{"zhang san", 12},
		{"li si", 30},
		{"wang wu", 52},
		{"zhao liu", 26},
	}

	fmt.Println(people)

	sort.Sort(PersonWrapper{people, func(p, q *Person) bool {
		return q.Age < p.Age // Age 递减排序
	}})

	fmt.Println(people)
	sort.Sort(PersonWrapper{people, func(p, q *Person) bool {
		return p.Name < q.Name // Name 递增排序
	}})

	fmt.Println(people)

}
