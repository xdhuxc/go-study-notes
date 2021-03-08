package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func worker2(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main1() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		go worker(i, &wg)
	}

	wg.Wait()
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("number: ", i)
		}(i)
	}
	wg.Wait()

	fmt.Println("lalalall")

	time.Sleep(time.Minute)
}
