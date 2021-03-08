package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"time"
)

const Dash = "-"
const Slash = "/"

const DashDateString = "2006-01-02"
const SlashDateString = "2006/1/2"

func main() {
	now := time.Now().Unix()
	fmt.Println(now)

	beforeSevenDays := time.Now().AddDate(0, 0, -7).Unix()

	fmt.Println(beforeSevenDays)

	fmt.Println(now - beforeSevenDays)

	// 格式化
	currentDateString := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(currentDateString)

	dashDate := "2019-08-10"
	slashDate := "2019/12/09"

	dashDateResult, err := time.Parse(DashDateString, dashDate)
	if err != nil {
		log.Errorln(err)
	} else {
		fmt.Println(dashDateResult)
	}

	slashDateResult, err := time.Parse(SlashDateString, slashDate)
	if err != nil {
		log.Errorln(err)
	} else {
		fmt.Println(slashDateResult)
	}
}
