package main

import "fmt"

type User struct {
	Hello
}

type Hello struct {
}

func (h Hello) H1() {
	fmt.Println("Hello 1")
}

type MyUser1 User
type MyUser2 = User

func (i MyUser1) m1() {
	fmt.Println("MyUser1.m1")
}
func (i MyUser2) m2() {
	fmt.Println("MyUser2.m2")
}

//Blog:www.flysnow.org
//Wechat:flysnow_org
func main() {
	var i1 MyUser1
	var i2 MyUser2
	i1.m1()
	i2.m2()
	i1.Hello.H1()
}
