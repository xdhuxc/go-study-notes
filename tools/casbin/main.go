package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/casbin/casbin"
	_ "github.com/go-sql-driver/mysql"

	gormadapter "github.com/casbin/gorm-adapter"
)

func main() {

	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"devops",
		"abcdsf",
		"10.1.9.16",
		"aye")
	a, err := gormadapter.NewAdapter("mysql", uri, true)
	if err != nil {
		log.Fatalln(err)
	}
	e, err := casbin.NewEnforcer("/Users/wanghuan/GolandProjects/GoPath/src/github.com/xdhuxc/go-study-notes/tools/casbin/auth_model.conf", a)

	if err != nil {
		log.Fatalln(err)
	}

	// 在运行时打开日志
	e.EnableLog(true)

	filter := gormadapter.Filter{
		PType: []string{"p"},
		V0:    []string{"admin"},
	}

	_ = e.LoadFilteredPolicy(filter)

	policies := e.GetPolicy()
	fmt.Println(len(policies))

}
