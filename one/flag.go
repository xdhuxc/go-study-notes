package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
)

var password string

func init() {
	flag.StringVar(&password, "password", "", "the password to authenticate with (only used with vault and etcd backends)")
}

func main() {

	flag.Parse()

	log.Println(password)

	if password == "qZ2Xis#dsdg#Gi#LSf0b$HBL#LDLUG5uUGr8v0" {
		log.Println(true)
	}

}
