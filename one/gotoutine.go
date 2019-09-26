package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}

	fmt.Println(id, x)
}

func goFunc() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			sum(id)
		}(i)
	}

	wg.Wait()
}

func goFuncDefer() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	wg.Wait()
}

func main() {

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i == 5 {
				runtime.Gosched()
			}
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Hello World")
	}()

	wg.Wait()
}
