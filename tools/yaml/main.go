package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/xdhuxc/chitchat/src/models"
)

func main() {
	var conf models.Configuration
	path := "conf.test.yml"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logrus.Fatalln("the configuration file does not exists", err)

	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Fatalln("Can not open configuration file", err)
	}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		logrus.Fatalln("Unmarshal yaml file error", err)
	}
	fmt.Println(conf.Server)

}
