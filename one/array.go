package main

import "fmt"

func testArray(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}

func main() {
	// 一维数组
	a1 := [3]int{1, 2}
	a2 := [...]int{1, 2, 3, 4, 5}
	a3 := [5]int{2: 100, 4: 300}

	a4 := [...]struct {
		name string
		age  uint8
	}{
		{"erha", 23},
		{"xdhuxc", 22},
	}

	// 多维数组
	a5 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	a6 := [...][2]int{{1, 1}, {2, 2}, {3, 3}}

	fmt.Println(a1, a2, a3, a4, a5, a6)

	a7 := [2]int{}
	fmt.Printf("a: %p\n", &a7)

	// 数组传递会复制该数组
	testArray(a7)
	fmt.Println(a7)

	fmt.Printf("length is %d, capacity is %d", len(a7), cap(a7))

}
