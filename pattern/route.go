package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	cm "gitlab.ushareit.me/sgt/scmp-common/src/models"
)

type Receiver struct {
	ID         int64     `json:"id" gorm:"id"`
	Name       string    `json:"name" gorm:"name"`
	GroupID    int       `json:"group_id" gorm:"group_id"`
	GroupName  string    `json:"group_name" gorm:"-"`
	Type       string    `json:"type" gorm:"type"`
	Resolved   bool      `json:"resolved" gorm:"resolved"`
	URL        string    `json:"url" gorm:"url"`
	CreateTime time.Time `json:"create_time" gorm:"create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"update_time"`
}

func (r *Receiver) TableName() string {
	return "sgt_hawkeye_receiver"
}

func (r *Receiver) String() string {
	if dataInBytes, err := json.Marshal(r); err == nil {
		return string(dataInBytes)
	}

	return ""
}

type Group struct {
	ID          int       `json:"id"`
	Name        string    `json:"group_name"`
	Description string    `json:"description"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
}

func (g *Group) TableName() string {
	return "sgt_hawkeye_group"
}

func (g *Group) Validate() error {
	if g.Name == "" {
		return errors.New("group name is empty")
	}
	return nil
}

type AlertGroup struct {
	Id          int     `json:"id" gorm:"column:id"`
	Name        string  `json:"name" gorm:"column:name"`
	Description string  `json:"description" gorm:"column:description"`
	Resolved    bool    `json:"resolved" gorm:"column:resolved"`
	DingtalkUrl string  `json:"dingtalk_url" gorm:"column:dingtalk_url"`
	GroupEmail  string  `json:"group_email" gorm:"column:group_email"`
	Tags        cm.JSON `json:"tags" gorm:"column:tags"`
	Members     cm.JSON `json:"members,omitempty" gorm:"-"`
}

func (ag *AlertGroup) TableName() string {
	return "sgt_scheduler_group"
}

func (ag *AlertGroup) Validate() error {
	if govalidator.IsNull(govalidator.Trim(ag.Name, "")) {
		return errors.New("group name can not be empty")
	}
	if govalidator.IsNull(govalidator.Trim(ag.DingtalkUrl, "")) {
		return errors.New("dingtalk url can not be empty")
	} else {
		if !govalidator.IsURL(ag.DingtalkUrl) {
			return errors.New("invalid dingtalk url")
		}
	}
	if govalidator.IsNull(govalidator.Trim(ag.GroupEmail, "")) {
		return errors.New("group email can not be empty")
	} else {
		if !govalidator.IsEmail(ag.GroupEmail) {
			return errors.New("invalid email address")
		}
	}

	return nil
}

var prodDBClient *gorm.DB
var testDBClient *gorm.DB

func init() {
	prodURI := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"devops",
		"iSfqaaatlzirpxh^rdkhykedMph3wd6i",
		"10.21.29.136:3306",
		"hawkeye")
	prodDBClient, _ = gorm.Open("mysql", prodURI)
	prodDBClient.Debug()

	testURI := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"Shareit@2018",
		"159.138.87.227:3306",
		"hawkeye")
	testDBClient, _ := gorm.Open("mysql", testURI)
	testDBClient.Debug()
}

type Route struct {
	ID             int       `json:"id" gorm:"id"`
	Name           string    `json:"name" gorm:"name"`
	GroupID        int       `json:"gid" gorm:"group_id"`
	AlertGroupID   int64     `json:"alert_group_id" gorm:"alert_group_id"`
	AlertGroupName string    `json:"alert_group_name" gorm:"-"`
	Match          cm.JSON   `json:"match,omitempty" gorm:"match" sql:"type:json"`
	MatchRE        cm.JSON   `json:"match_re,omitempty" gorm:"match_re" sql:"type:json"`
	GroupWait      int       `json:"group_wait" gorm:"group_wait"`
	GroupInterval  int       `json:"group_interval" gorm:"group_interval"`
	RepeatInterval int       `json:"repeat_interval" gorm:"repeat_interval"`
	CreateTime     time.Time `json:"create_time" gorm:"create_time"`
	UpdateTime     time.Time `json:"update_time" gorm:"update_time"`
}

func (r *Route) TableName() string {
	return "sgt_hawkeye_route"
}

func main1() {
	var a *Route

	fmt.Println(a)
}

func TransformReceiver() {
	var receivers []Receiver
	if err := prodDBClient.Model(&Receiver{}).Find(&receivers).Error; err != nil {
		logrus.Error(err)
		return
	}
	length := len(receivers)
	// ags := make([]AlertGroup, length)
	for i := 0; i < length; i++ {
		var g Group
		if err := prodDBClient.Model(&Group{}).Where("id = ?", receivers[i].GroupID).First(&g).Error; err != nil {
			logrus.Error(err)
			continue
		}
		receivers[i].GroupName = g.Name

		var ag AlertGroup
		ag.Name = receivers[i].Name
		ag.Description = strings.ReplaceAll(receivers[i].Name, "_", " ")
		ag.Resolved = receivers[i].Resolved
		if receivers[i].Type == "dingtalk" {
			ag.DingtalkUrl = receivers[i].URL
			ag.GroupEmail = strings.ToLower(receivers[i].GroupName) + "@ushareit.com"
		} else if receivers[i].Type == "email" {
			ag.GroupEmail = receivers[i].URL
			ag.DingtalkUrl = "https://oapi.dingtalk.com/robot/send?access_token"
		}

		if err := ag.Validate(); err != nil {
			logrus.Error(err)
			continue
		}

		if err := testDBClient.Model(&AlertGroup{}).Create(&ag).Error; err != nil {
			logrus.Error(err)
			continue
		}
	}
}
