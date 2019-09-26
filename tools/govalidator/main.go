package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
)

func main() {
	s := "   "
	fmt.Println(len(s))

	fmt.Println(len(govalidator.Trim(s, "")))
}
