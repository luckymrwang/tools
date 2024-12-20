package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"reflect"
	"regexp"
	"runtime/debug"
	"strconv"
	"testing"
	"unicode/utf8"
)

func TestEcho(t *testing.T) {
	dService := GetDemoService(nil)
	dService.Echo()
}

func TestGet(t *testing.T) {
	_, cidr, err := net.ParseCIDR("10.233.63.0/18")
	fmt.Println(cidr.String())

	s := "xx1x王"
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println(len(s))

	_, err = strconv.ParseInt("10000000000", 10, 64)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSlice(t *testing.T) {
	a := []string{"a", "b", "c"}
	for _, s := range a {
		fmt.Println(&s)
	}
	fmt.Println(a)
}

type AS struct {
	AA string
	BB string
}

func TestMap(t *testing.T) {
	var a []string
	data := map[string]interface{}{
		"a": a,
	}
	buf, err := json.Marshal(&data)
	log.Println(string(buf), err)
	data = map[string]interface{}{
		"b": make([]string, 0),
	}
	buf, err = json.Marshal(&data)
	log.Println(string(buf), err)
}

func TestAS(t *testing.T) {
	as := new(AS)
	fmt.Println(*as)
	ass := AS{}
	fmt.Println(ass)

	f := float64(30.33)
	prt()
	fmt.Printf("%.f", f)
}

func prt() {
	debug.PrintStack()
}

func TestCopy(t *testing.T) {
	map1 := map[string]string{"1": "1", "2": "2", "3": "3"}
	mp2 := map1
	fmt.Printf("[old] address: %p, values: %v\n", map1, map1)
	fmt.Printf("[new] address: %p, values: %v\n", mp2, mp2)
	t.Logf("修改map1中的一个值后")
	map1["1"] = "100"
	fmt.Printf("[old] address: %p, values: %v\n", map1, map1)
	fmt.Printf("[new] address: %p, values: %v\n", mp2, mp2)
}

func TestMatch(t *testing.T) {
	const qnameCharFmt string = "[A-Za-z0-9]"
	const qnameExtCharFmt string = "[-A-Za-z0-9_.]"
	const qualifiedNameFmt string = "(" + qnameCharFmt + qnameExtCharFmt + "*)?" + qnameCharFmt
	const qualifiedNameErrMsg string = "must consist of alphanumeric characters, '-', '_' or '.', and must start and end with an alphanumeric character"
	const qualifiedNameMaxLength int = 63

	var qualifiedNameRegexp = regexp.MustCompile("^" + qualifiedNameFmt + "$")
	name := `a/b-d`
	if !qualifiedNameRegexp.MatchString(name) {
		fmt.Println("not match")
	} else {
		fmt.Println("match")
	}
}

func TestSymmetricDifference(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []int
		slice2   []int
		expected []int
	}{
		{"empty slices", []int{}, []int{}, []int{}},
		{"no difference", []int{1, 2, 3}, []int{1, 2, 3}, []int{}},
		{"some difference1", []int{1, 2, 3}, []int{2, 3, 4}, []int{1, 4}},
		{"some difference2", []int{1, 2}, []int{2, 3, 4}, []int{1, 3, 4}},
		{"some difference3", []int{1}, []int{1, 3, 4}, []int{3, 4}},
		{"some difference4", []int{1, 2, 3}, []int{2, 3}, []int{1}},
		{"all difference", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := SymmetricDifference(tt.slice1, tt.slice2)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("difference(%v, %v) = %v, want %v", tt.slice1, tt.slice2, actual, tt.expected)
			}
		})
	}
}
