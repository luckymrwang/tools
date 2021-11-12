package models

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestDH_Corp_OrderPager(t *testing.T) {
	_, list, err := new(DH_Corp).OrderPager(map[string]interface{}{"status": -1}, 1, 1, "create_time", "desc")
	if err != nil {
		t.Error(err)
		return
	}
	for _, dhc := range list {
		tx := DB.Begin()
		fmt.Println(dhc.Name)
		err := new(DH_Corp).Update(dhc, "hello", tx)
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
			continue
		}
		err2 := new(DH_Corp).Delete(dhc, tx)
		if err2 != nil {
			tx.Rollback()
			fmt.Println(err2)
			continue
		}
		tx.Commit()
	}
}

func TestTM(t *testing.T) {
	dh := &DH_Corp{}
	err := new(DH_Corp).TM(dh)
	if err != nil {
		t.Error(err)
	}
	return
}

type Jac struct {
	F1 float64 `json:"f_1"`
	F2 float64 `json:"f_2"`
}

func (j *Jac) Echo() {
	fmt.Println("hello")
}

func TestGet(t *testing.T) {
	instance := &Jac{F1: 1000000000, F2: 20.2}
	v := reflect.ValueOf(instance)
	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.Float64 {
			val := v.Field(i).Float()
			if val == float64(int(val)) {
				fmt.Printf("%v:%f is int", v.Type().Field(i).Name, v.Field(i).Float())
			} else {
				fmt.Printf("%v:%f is float", v.Type().Field(i).Name, v.Field(i).Float())
			}
		}
	}
}

func TestPrifix(t *testing.T) {
	a := "1-2"
	b := "1-2-3"
	c := "1-21-3"
	if strings.HasPrefix(b, a+"-") {
		fmt.Printf("ok")
	}
	if strings.HasPrefix(c, a+"-") {
		t.Error("false")
	}
}
