package main

import (
	"fmt"
	"math/rand"
	"time"
)

func eat() chan string {
	out := make(chan string)
	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		out <- "do eat"
		fmt.Println("i eat and drink")
		close(out)
	}()

	return out
}

func main() {
	i := 63102
	fmt.Println(i / 1000)

	/*
		for {
			ec := eat()
			sleep := time.NewTimer(time.Second * 3)
			select {
			case s := <- ec:
				fmt.Println(s)
			case <- sleep.C:
				fmt.Println("Time to sleep")
			default:
				fmt.Println("Beat Doudou")
			}
		}

	*/
}
