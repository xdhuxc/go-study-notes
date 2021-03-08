package main

import (
	"fmt"
	"github.com/panjf2000/ants"
	"github.com/sirupsen/logrus"
	"sync/atomic"
	"time"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func Hello() {
	fmt.Println("hello world")
}

func Sum(a, b int) int {
	return a + b
}

func main() {
	defer ants.Release()

	pool, err := ants.NewPool(10)
	if err != nil {
		logrus.Error(err)
		return
	}

	fmt.Println("协程池容量：", pool.Cap())
	fmt.Println("空闲协程池容量：", pool.Free())
	fmt.Println("运行协程池容量：", pool.Running())

	_ = pool.Submit(Hello)
	fmt.Println("协程池容量：", pool.Cap())
	fmt.Println("空闲协程池容量：", pool.Free())
	fmt.Println("运行协程池容量：", pool.Running())

	_ = pool.Submit(func() {
		x := Sum(2, 3)
		fmt.Println(x)
	})

	fmt.Println("协程池容量：", pool.Cap())
	fmt.Println("空闲协程池容量：", pool.Free())
	fmt.Println("运行协程池容量：", pool.Running())

	time.Sleep(1 * time.Minute)
}
