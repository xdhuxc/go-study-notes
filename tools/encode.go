package main

import (
	"fmt"

	"github.com/axgle/mahonia"
)

func main() {

	decoder := mahonia.NewDecoder("GBK")
	// hello, 涓栫晫
	fmt.Println(decoder.ConvertString("hello, 世界"))

	dataInBytes := []byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1}

	// 你好，世界！
	fmt.Println(decoder.ConvertString(string(dataInBytes)))
}
