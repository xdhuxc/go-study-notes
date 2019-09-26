package main

import "fmt"

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	B int64 = 1 << (10 * iota)
	KB
	MB
	GB
	TB
)

type Color int

const (
	Black Color = iota
	Red
	Blue
)

func test(c Color) {

}

func main() {
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)

	c := Black
	test(c)

	/**
	x := 1
	test(x) // can not use type int as type Color
	*/

	test(1)
}
