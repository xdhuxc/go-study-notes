package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()

	glog.Infoln("This is a info log for glog")
	glog.Infof("This is a info format log for %s", "glog")

	glog.Warningln("This is a warning log for glog")
	glog.Warningf("This is a warning format log for %s", "glog")

	glog.Errorln("This is a error log for glog")
	glog.Errorf("This is a error format log for %s", "glog")

	/*
		glog.Fatalln("This is a fatal log for glog")
		glog.Fatalf("This is a fatal format log for %s", "glog")
	*/

	var a *string
	var b *string

	c := "abc"

	a = &c
	b = &c

	fmt.Println(a == &c)
	fmt.Println(b == &c)

	glog.Flush()
}
