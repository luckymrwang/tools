package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name   string
	Weight int
}

type PersonSlice []Person

func (s PersonSlice) Len() int      { return len(s) }
func (s PersonSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByName struct{ PersonSlice } // 将 PersonSlice 包装起来到 ByName 中

func (s ByName) Less(i, j int) bool { return s.PersonSlice[i].Name < s.PersonSlice[j].Name } // 将 Less 绑定到 ByName 上

type ByWeight struct{ PersonSlice }   // 将 PersonSlice 包装起来到 ByWeight 中
func (s ByWeight) Less(i, j int) bool { return s.PersonSlice[i].Weight < s.PersonSlice[j].Weight } // 将 Less 绑定到 ByWeight 上

func main() {
	s := []Person{
		{"apple", 12},
		{"pear", 20},
		{"banana", 50},
		{"orange", 87},
		{"hello", 34},
		{"world", 43},
	}

	sort.Sort(ByWeight{s})
	fmt.Println("People by weight:")
	printPeople(s)

	sort.Sort(ByName{s})
	fmt.Println("\nPeople by name:")
	printPeople(s)

}

func printPeople(s []Person) {
	for _, o := range s {
		fmt.Printf("%-8s (%v)\n", o.Name, o.Weight)
	}
}
