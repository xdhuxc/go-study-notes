package main

import (
	"fmt"
	"strconv"

	"github.com/golang/glog"
)

func main() {

	u := "笔记本"
	us := []rune(u)

	us[1] = '查'
	fmt.Println(len(us))
	fmt.Println(string(us))

	glog.Errorf("the response is not ok, the status code is %v", strconv.Itoa(500))

}
