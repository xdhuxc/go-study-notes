package main

import (
	"fmt"
	"math"
)

func main() {

	r := 0.75*(0.05*0.1+0.04*0.1+0.03*0.1+0.02*0.1+0.01*0.1) - 0.25*(0.04*0.1+0.03*0.1+0.02*0.1+0.01*0.1)

	fmt.Println(r)

	count := 12
	pageSize := 5
	result := int64(math.Ceil(float64(count) / float64(pageSize)))
	fmt.Println(result)
}
