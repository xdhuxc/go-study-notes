package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id int `json:"id" gorm:"id"`

	CreateTime time.Time `json:"create_time" gorm:"create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"update_time"`

	XdhuxcTime time.Time `json:"xdhuxc_time" gorm:"xdhuxc_time"`

	AbcTime time.Time `json:"abc_time" gorm:"column:abc_time;type:timestamp"`
}

func (u User) String() string {
	if dataInBytes, err := json.Marshal(u); err == nil {
		return string(dataInBytes)
	}

	return ""
}

func (u *User) TableName() string {
	return "users"
}

func main() {
	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"wh19940423",
		"localhost:3306",
		"stock")
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		log.Errorln(err)
		return
	}
	db.Debug()
	db.LogMode(true)

	u := User{
		Id:         24,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		XdhuxcTime: time.Now(),
		AbcTime:    time.Now(),
	}

	fmt.Println(u.String())

	if err := db.Model(&User{}).Create(&u).Error; err != nil {
		log.Println(err)
	}

	var x User
	if err := db.Model(&User{}).Where("id = ?", 4).First(&x).Error; err != nil {
		log.Println(err)
	}
	fmt.Println(x.String())
}
