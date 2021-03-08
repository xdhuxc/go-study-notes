package main

import (
	"fmt"
	"unsafe"

	json "github.com/json-iterator/go"
)

//var json = jsoniter.ConfigCompatibleWithStandardLibrary

type MyStruct1 struct {
	i    int    `json:"i"`
	Name string `json:"name"`
}

type MyStruct2 struct {
	i int
	j int
}

func myFunction2(ms *MyStruct2) {
	ptr := unsafe.Pointer(ms)
	for i := 0; i < 2; i++ {
		c := (*int)(unsafe.Pointer(uintptr(ptr) + uintptr(8*i)))
		*c += i + 1
		fmt.Printf("[%p] %d\n", c, *c)
	}
}

func myFunction1(a MyStruct1, b *MyStruct1) {
	a.i = 31
	b.i = 41
	fmt.Printf("in my_function - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
}

func myFunction(i *int, arr [2]int) {
	*i = 100
	arr[1] = 100
	fmt.Printf("in my_funciton - i=(%d, %p) arr=(%v, %p)\n", *i, &i, arr, &arr)
}

func main() {
	i := 30
	arr := [2]int{66, 77}
	fmt.Printf("before calling - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
	myFunction(&i, arr)
	fmt.Printf("after  calling - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)

	a := MyStruct1{i: 30}
	b := &MyStruct1{i: 40}
	fmt.Printf("before calling - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
	myFunction1(a, b)
	fmt.Printf("after calling  - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)

	x := &MyStruct2{i: 40, j: 50}
	myFunction2(x)
	fmt.Printf("[%p] %v\n", x, x)

	dataInBytes, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(dataInBytes))
}
