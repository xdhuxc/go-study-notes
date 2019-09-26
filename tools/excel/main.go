package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	log "github.com/sirupsen/logrus"
)

func main() {
	filePath := "/Users/wanghuan/Library/Mobile Documents/com~apple~Numbers/Documents/个人文件/2018年银行卡账目明细.xlsx"
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range xlsx.GetSheetMap() {
		fmt.Println(name)
		rows, err := xlsx.Rows(name)
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			cols, err := rows.Columns()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(cols)
			//for _, colCell := range cols {
			//	fmt.Println(colCell)
			//}
		}
	}
}
