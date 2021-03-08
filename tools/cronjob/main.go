package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

// https://www.jianshu.com/p/fd3dda663953
func main() {
	c := cron.New()

	c.AddFunc("*/1 * * * ?", func() {
		fmt.Println("每隔1分钟执行一次")
	})

	c.AddFunc("*/2 * * * ?", func() {
		fmt.Println("每隔2分钟执行一次")
	})

	c.AddFunc("@every 1m", func() {
		fmt.Println("every minute")
	})

	c.Start()

	c.AddFunc("@hourly", func() {
		fmt.Println("every hour")
	})
	for {
		select {
		case <-time.After(time.Minute):
			fmt.Println("一分钟过去了")
		}
	}
}
