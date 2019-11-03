package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Unix()
	fmt.Println(now)

	beforeSevenDays := time.Now().AddDate(0, 0, -7).Unix()

	fmt.Println(beforeSevenDays)

	fmt.Println(now - beforeSevenDays)
}
