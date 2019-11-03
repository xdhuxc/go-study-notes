package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// 计算 2 的 5 次方
	fmt.Println(math.Pow(2, 5))

	percent := "23.456"
	// 使用 strconv 包的 ParseFloat 方法将字符串转换为 float64
	result, _ := strconv.ParseFloat(percent, 10)
	// 使用 fmt 包的 Sprintf 方法实现保留小数点后两位
	fmt.Println(fmt.Sprintf("%.2f", result))

}
