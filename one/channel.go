package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Request struct {
	data []int
	ret  chan int
}

func NewRequest(data ...int) *Request {
	return &Request{data, make(chan int, 1)}
}

func Process(req *Request) {
	x := 0
	for _, i := range req.data {
		x += i
	}

	req.ret <- x
}

func main() {
	req := NewRequest(10, 20, 30)
	Process(req)
	fmt.Println(<-req.ret)
}

func test9() {
	w := make(chan bool)
	c := make(chan int, 2)

	go func() {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout")
		}

		w <- true
	}()

	// c <-1
	<-w
}

func test8() {
	var wg sync.WaitGroup
	quit := make(chan bool)

	for i := 0; i < 2; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			task := func() {
				fmt.Println(id, time.Now().Nanosecond())
				time.Sleep(time.Second)
			}

			for {
				select {
				case <-quit: // closed channel 不会阻塞，因此可用作退出通知
					return
				default:
					task() // 执行正常任务
				}
			}

		}(i)
	}

	time.Sleep(time.Second * 5)

	close(quit)

	wg.Wait()
}

func test7() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	semaphore := make(chan int, 1)
	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()

			semaphore <- 1

			for x := 0; x < 3; x++ {
				fmt.Println(id, x)
			}

			<-semaphore

		}(i)
	}

	wg.Wait()
}

func test6() {
	t := NewTest()
	fmt.Println(<-t)
}

func NewTest() chan int {
	c := make(chan int)
	rand.Seed(time.Now().UnixNano())

	go func() {
		time.Sleep(time.Second)
		c <- rand.Int()
	}()

	return c
}

func test5() {
	a, b := make(chan int, 3), make(chan int)

	go func() {
		v, ok, s := 0, false, ""
		for {
			select {
			case v, ok = <-a:
				s = "a"
			case v, ok = <-b:
				s = "b"
			}

			if ok {
				fmt.Println(s, v)
			} else {
				os.Exit(0)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		select {
		case a <- i:
		case b <- i:
		}
	}

	close(a)

	select {}
}

func test4() {
	c := make(chan int, 3)
	// 仅用于发送数据
	var send chan<- int = c
	// 仅用于接收数据
	var receive <-chan int = c

	send <- 1
	// Error: receive from send-only type chan<- int
	// <-send

	<-receive
	// Error: send to receive-only type <-chan int
	// receive <- 2
}

func test3() {
	// 缓冲区可以存储 3 个元素
	data := make(chan int, 3)
	exit := make(chan bool)

	// 在缓冲区未满前，不会阻塞
	data <- 1
	data <- 2
	data <- 3

	go func() {

		// 在缓冲区未空前，不会阻塞
		for {
			if d, ok := <-data; ok {
				fmt.Println(d, len(data), cap(data))
			} else {
				break
			}
		}

		exit <- true
	}()

	/**
	cap() 返回缓冲区大小
	len() 返回未被读取的缓冲元素数量
	*/
	fmt.Println(len(data), cap(data))
	data <- 4
	fmt.Println(len(data), cap(data))
	data <- 5

	close(data)

	<-exit
}

func test2() {
	// 缓冲区可以存储 3 个元素
	data := make(chan int, 3)
	exit := make(chan bool)

	// 在缓冲区未满前，不会阻塞
	data <- 1
	data <- 2
	data <- 3

	go func() {
		for d := range data { // 在缓冲区未空前，不会阻塞
			fmt.Println(d)
		}

		exit <- true
	}()

	data <- 4
	data <- 5

	close(data)

	<-exit
}

func test1() {
	// 数据交换队列
	data := make(chan int)
	// 退出通知
	exit := make(chan bool)

	go func() {
		for d := range data {
			fmt.Println(d)
		}

		fmt.Println("recv over.")
		// 发出退出通知
		exit <- true
	}()

	// 发送数据
	data <- 1
	data <- 2
	data <- 3

	// 关闭队列
	close(data)

	fmt.Println("send over...")
	// 等待退出通知
	<-exit
}
