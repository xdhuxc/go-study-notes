package main

import (
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

func main() {
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
