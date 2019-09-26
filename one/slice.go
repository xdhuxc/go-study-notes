package main

import "fmt"

func main() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6}
	s1 := data[1:4:5]

	fmt.Println(s1)

	data2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// cap = max - low
	fmt.Printf("the slice is %d, the length is %d, the capacity is %d\n", data2[:6:8], len(data2[:6:8]), cap(data2[:6:8]))
	fmt.Printf("the slice is %d, the length is %d, the capacity is %d\n", data2[5:], len(data2[5:]), cap(data2[5:]))
	fmt.Printf("the slice is %d, the length is %d, the capacity is %d\n", data2[:3], len(data2[:3]), cap(data2[:3]))
	fmt.Printf("the slice is %d, the length is %d, the capacity is %d\n", data2[:], len(data2[:]), cap(data2[:]))

	s2 := []int{0, 1, 2, 3, 8: 100}
	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 6, 8)
	fmt.Println(s3, len(s3), cap(s3))

	s4 := make([]int, 6)
	fmt.Println(s4, len(s4), cap(s4))

	data3 := [...]int{0, 1, 2, 3, 4, 10: 0}
	s5 := data3[:2:3]

	s5 = append(s5, 200)
	fmt.Println(data3, s5)

	fmt.Println(&s5[0], &data3[0])

	s6 := make([]int, 0, 1)
	c := cap(s6)

	for i := 0; i < 50; i++ {
		s6 = append(s6, i)
		if n := cap(s6); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}

	data4 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s7 := data4[8:]
	s8 := data[:5]
	fmt.Println(s7, s8)

	copy(s8, s7)

	fmt.Println(s8, data4)

	fmt.Println("===================")

	abc := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 20, 11}
	fmt.Println(abc[5:])
}
