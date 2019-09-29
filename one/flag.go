package main

import (
	"flag"
	"fmt"
	"os"
)

var password string

func init() {
	flag.StringVar(&password, "password", "", "the password to authenticate with (only used with vault and etcd backends)")
}

func main() {
	fmt.Printf("the args is %s\n", os.Args)

	flag.Parse()

	fmt.Printf("the password is %s\n", password)

	if password == "qZ2XisdsdgGiLSf0b$HBL#LDLUG5uUGr8v0" {
		fmt.Println(true)
	}
}
