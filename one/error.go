package main

import "fmt"

// 使用 panic 抛出异常，recover 捕获异常，panic 之后的代码不再执行
func error1() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err.(string))
		}
	}()

	panic("panic error")

	fmt.Println("after panic")
}

func error2() {
	/**
	延迟调用中引发的错误，可被后续延迟调用捕获，但是仅最后一个错误可被捕获
	*/
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("first defer panic")
	}()

	defer func() {
		panic("second defer panic")
	}()

	defer func() {
		panic("third defer panic")
	}()

	panic("test panic")
}

func main() {
	/**
	在主进程中，即使一个函数发生了 panic，不影响其他函数的 panic。
	*/
	error1()

	error2()
}
