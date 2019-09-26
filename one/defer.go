package main

import (
	"fmt"
	"os"
)

func add(x int, y int) (z int) {

	defer func() {
		z = z + 200
	}()

	z = x + y
	return z
}

func testDefer1() error {
	f, err := os.Create("abcd.txt")
	if err != nil {
		return err
	}

	defer f.Close()

	num, err := f.WriteString("Hello World!")
	fmt.Println(num)
	return err
}

func testDefer2(x int) {
	/**
	多个 defer 注册，按 FILO 次序执行，即使函数或某个延迟调用发生错误，这些调用依旧会被执行
	*/
	defer fmt.Println("a")
	defer fmt.Println("b")

	defer func() {
		fmt.Println(100 / x)
	}()

	defer fmt.Println("c")

	fmt.Println("test defer 2")
}

func testDefer3() {
	/**
	延迟调用参数在注册时求值或复制，可用指针或闭包 "延迟读取"
	*/
	x, y := 10, 20

	defer func(i int) {
		fmt.Println("defer: ", i, y) // y 闭包引用，y 值为120
	}(x) // x 被复制，x 值为 10

	x += 10
	y += 100

	fmt.Println("x = ", x, "y = ", y)
}

func main() {

	err := testDefer1()
	if err != nil {
		fmt.Println(err)
	}

	testDefer2(5)

	testDefer3()

}
