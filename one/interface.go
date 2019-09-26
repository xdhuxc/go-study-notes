package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Stringer interface {
	String() string
}

type Printer interface {
	Stringer
	Print()
}

type User struct {
	id   int
	name string
}

func (u *User) String() string {
	return fmt.Sprintf("User: %p, %v", u, &u)
}

func (u *User) Print() {
	fmt.Println(u.String())
}

func PrintInterface(v interface{}) {
	fmt.Printf("%T: %v\n", v, v)
}

var a interface{} = nil
var b interface{} = (*int)(nil)

type iface struct {
	itab uintptr
	data uintptr
}

func main() {
	var t Printer = &User{1, "xdhuxc"}

	t.String()
	fmt.Println("==============")
	t.Print()

	PrintInterface(1)
	PrintInterface("abcd")
	PrintInterface(1.234)

	fmt.Println("==============")
	u := User{1, "xdhuxc"}
	var vi, pi interface{} = u, &u
	// vi.(User).name = "xxx" cannot assign to vi.(User).name
	pi.(*User).name = "yyy"

	fmt.Printf("%v\n", vi.(User))
	fmt.Printf("%v\n", pi.(*User))

	fmt.Println("==============")
	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))

	fmt.Println(a == nil, ia)
	fmt.Println(b == nil, ib, reflect.ValueOf(b).IsNil())

	fmt.Println("==============")
	var o interface{} = &User{1, "xdhuxc"}

	if i, ok := o.(fmt.Stringer); ok {
		fmt.Println(i)
	}
	u2 := o.(*User)
	fmt.Println(u2)
}
