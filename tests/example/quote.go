package main

type B struct {
	Id int
}

func New() B {
	return B{}
}

func New2() *B {
	return &B{}
}

func (b *B) Hello() {
	return
}

func (b B) World() {
	return
}

func main() {
	// 方法的接收器为 *T 类型
	New().Hello() // 编译不通过

	b1 := New()
	b1.Hello() // 编译通过

	b2 := B{}
	b2.Hello() // 编译通过

	(B{}).Hello() // 编译不通过
	B{}.Hello()   // 编译不通过

	New2().Hello() // 编译通过

	b3 := New2()
	b3.Hello() // 编译通过

	b4 := &B{} // 编译通过
	b4.Hello() // 编译通过

	(&B{}).Hello() // 编译通过

	// 方法的接收器为 T 类型
	New().World() // 编译通过

	b5 := New()
	b5.World() // 编译通过

	b6 := B{}
	b6.World() // 编译通过

	(B{}).World() // 编译通过
	B{}.World()   // 编译通过

	New2().World() // 编译通过

	b7 := New2()
	b7.World() // 编译通过

	b8 := &B{} // 编译通过
	b8.World() // 编译通过

	(&B{}).World() // 编译通过
}
