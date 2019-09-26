package main

import (
	"fmt"
	"unsafe"
)

var x, y, z int
var s, n = "abc", 123

var (
	a int
	b float32
)

const (
	c = "abc"
	d
)

const (
	e = "abc"
	f = len(e)
	g = unsafe.Sizeof(b)
)

const (
	h byte = 100
	// i int = 1e20 constant 100000000000000000000 overflows int64
)

func main() {
	fmt.Println(&s, &n)
	// 重新定义同名变量
	n, s := 0x1234, "Hello, World!"
	fmt.Println(x, s, n)
	fmt.Println(&s, &n)
	fmt.Println(h)

}
