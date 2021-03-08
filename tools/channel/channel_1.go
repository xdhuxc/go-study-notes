package main

import "fmt"

func main1() {
	rc := make(chan int, 10)
	wc := make(chan int, 10)

	// x := 10
	y := 100
	select {
	case x := <-rc:
		fmt.Printf("read %d \n", x)
	case wc <- y:
		fmt.Printf("write %d \n", y)
	default:
		fmt.Println("do default")
	}

}
