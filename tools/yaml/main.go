package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Server Server `yaml:"Server"`
}

type Server struct {
	Address      string `yaml:"Address"`
	ReadTimeout  int64  `yaml:"ReadTimeout"`
	WriteTimeout int64  `yaml:"WriteTimeout"`
	Static       string `yaml:"Static"`
	Password     string `yaml:"Password"`
}

func main() {
	var conf Configuration
	path := "/Users/wanghuan/GolandProjects/GoPath/src/github.com/xdhuxc/go-study-notes/tools/yaml/conf.test.yml"

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
