package main

import "fmt"

func main() {
	defer fmt.Println("in main")
	defer fmt.Println("after in main")
	defer func() {
		defer func() {
			fmt.Println("panic again and again 2")
			panic("panic again and again")
		}()
		fmt.Println("panic again 1")
		panic("panic again")
	}()
	defer func() {
		fmt.Println("second defer")
	}()

	panic("panic once")

	fmt.Println("after panic")
}
