package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	yawg := wg
	fmt.Println(wg, yawg)

	o := &sync.Once{}
	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Println("only once")
		})
		o.Do(func() {
			fmt.Println("only once")
		})
		o.Do(func() {
			fmt.Println("only once")
		})
	}
}
