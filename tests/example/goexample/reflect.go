import (
	"fmt"
)

type TypeOne struct {
}

func (t *TypeOne) FuncOne() {
	fmt.Println("FuncOne")
}

func (t *TypeOne) FuncTwo(name string) {
	fmt.Println("Hello", name)
}