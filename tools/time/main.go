package main

import (
	"fmt"
	"time"
)

func main() {
	d := time.Duration(time.Second * 2)
	timer := time.NewTimer(d)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			fmt.Println(time.Now().String())
			timer.Reset(d)
		}
	}
}

func main2() {
	d := time.Duration(time.Second * 2)

	// ticker 只要定义完成，从程序启动时刻开始，不需要任何其他的操作，每隔固定时间触发操作
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now().String())
		}
	}
}

func main1() {
	timeout := 10
	select {
	case <-time.After(time.Duration(timeout) * time.Second):
		fmt.Println("Time out ")
	}

	current := time.Now()
	fmt.Println(current.String())
	for {
		if time.Since(current) >= (time.Duration(timeout) * time.Second) {
			fmt.Println("it is end")
			fmt.Println(time.Now().String())
			return
		}
	}
}
