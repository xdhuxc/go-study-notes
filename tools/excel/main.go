package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// 银行卡账目明细表
type BankTransactionDetails struct {
	ID                    int64     `json:"id" gorm:"id"`
	Date                  time.Time `json:"date" gorm:"date"`
	Type                  string    `json:"type" gorm:"type"`
	BankCard              string    `json:"bankCard" gorm:"bank_card"`
	Income                float64   `json:"income" gorm:"income"`
	Expenditure           float64   `json:"expenditure" gorm:"expenditure"`
	Balance               float64   `json:"balance" gorm:"balance"`
	Remark                string    `json:"remark" gorm:"remark"`
	AdditionalDescription string    `json:"additionalDescription" gorm:"additional_description"`
	CreateTime            time.Time `json:"createTime" gorm:"create_time"`
	UpdateTime            time.Time `json:"updateTime" gorm:"update_time"`
}

func (btd *BankTransactionDetails) TableName() string {
	return "ams-bank_transaction_details"
}

func (btd *BankTransactionDetails) String() string {
	if dataInBytes, err := json.Marshal(&btd); err == nil {
		return string(dataInBytes)
	}

	return ""
}

func main() {
	fileFullPath := "/Users/wanghuan/Library/Mobile Documents/com~apple~Numbers/Documents/个人文件/2019年银行卡账目明细.xlsx"
	f, err := excelize.OpenFile(fileFullPath)
	if err != nil {
		log.Fatal(err)
	}

	for sheetIndex, sheetName := range f.GetSheetMap() {
		fmt.Println(sheetIndex, "--->", sheetName)
		rows, err := f.Rows(sheetName)
		if err != nil {
			log.Errorln(err)
			continue
		}

		for rows.Next() {
			// 获取一行
			var btd BankTransactionDetails

			/*
				if sheetName == "中国银行卡" {
					btd.BankCard = "BOC"
				} else if sheetName == "" {
					btd.BankCard = ""
				} else if sheetName == "" {
					btd.BankCard = ""
				} else if sheetName == "" {
					btd.BankCard = ""
				} else if sheetName == "" {
					btd.BankCard = ""
				} else if sheetName == "" {
					btd.BankCard = ""
				}
			*/
			btd.BankCard = sheetName
			cols, err := rows.Columns()
			if err != nil {
				log.Errorln(err)
				continue
			}
			// 获取行数组中的值
			if len(cols) > 7 {
				fmt.Println(cols)
			}
		}

	}

}

func main1() {
	users := []User{
		{
			ID:      1,
			Name:    "孔子",
			Address: "山东曲阜",
		},
		{
			ID:      2,
			Name:    "牛顿",
			Address: "英国伦敦",
		},
		{
			ID:      3,
			Name:    "凯撒",
			Address: "罗马",
		},
	}

	f := excelize.NewFile()
	sheetName := "人物信息"
	sheet := f.NewSheet(sheetName)
	f.SetActiveSheet(sheet)

	// 合并单元格
	err := f.MergeCell(sheetName, "A1", "A5")

	err = f.SaveAs("./k8s.xlsx")
	if err != nil {
		log.Errorln(err)
	}

	header := map[string]interface{}{"A1": "编号", "B1": "姓名", "C1": "地址"}
	data := generateData(users, header)
	err = Write2File("历史人物", sheetName, data, true)
	if err != nil {
		log.Errorln(err)
	}

	Read("/Users/wanghuan/GolandProjects/GoPath/src/github.com/xdhuxc/go-study-notes/历史人物_2019-10-22_.xlsx")
}

func generateData(users []User, header map[string]interface{}) []map[string]interface{} {
	var maps []map[string]interface{}
	var m map[string]interface{}

	maps = append(maps, header)
	rowCount := 1
	for _, user := range users {
		index := strconv.Itoa(rowCount)
		m = map[string]interface{}{
			"A" + index: user.ID,
			"B" + index: user.Name,
			"C" + index: user.Address,
		}

		maps = append(maps, m)
		rowCount = rowCount + 1
	}

	return maps
}

func Read(fileFullPath string) {
	f, err := excelize.OpenFile(fileFullPath)
	if err != nil {
		log.Fatal(err)
	}

	for sheetIndex, sheetName := range f.GetSheetMap() {
		fmt.Println(sheetIndex, "--->", sheetName)
		rows, err := f.Rows(sheetName)
		if err != nil {
			log.Errorln(err)
			continue
		}
		for rows.Next() {
			// 获取一行
			cols, err := rows.Columns()
			if err != nil {
				log.Errorln(err)
				continue
			}
			// 获取行数组中的值
			for _, colCell := range cols {
				fmt.Println(colCell)
			}
		}
	}
}

// 写数据到 excel 中，每个 map[string]interface{} 为一行数据，键为A1，B2，C3等等
func Write2File(fileName string, sheetName string, data []map[string]interface{}, dateContaining bool) error {
	f := excelize.NewFile()

	sheet := f.NewSheet(sheetName)
	f.SetActiveSheet(sheet)

	for _, item := range data {
		for k, v := range item {
			err := f.SetCellValue(sheetName, k, v)
			if err != nil {
				log.Errorln(err)
				continue
			}
		}
	}

	// 删除 Sheet1
	f.DeleteSheet("Sheet1")

	if dateContaining {
		return f.SaveAs(fileName + "_" + time.Now().Format("2006-01-02") + "_" + ".xlsx")
	} else {
		return f.SaveAs(fileName + ".xlsx")
	}
}
