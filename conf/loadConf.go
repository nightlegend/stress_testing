package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"time"
)

type Configure struct {
	Host          string        `yaml:"host"`
	Port          string        `yaml:"port"`
	Times         time.Duration `yaml:"times"`
	Rate          int64         `yaml:"rate"`
	ResultPath    string        `yaml:"resultPath"`
	TestResulName string        `yaml:"testResulName"`
	RequestType   string        `yaml:"requestType"`
	RequestPath   string        `yaml:"requestPath"`
	RequestData   string        `yaml:"requestData"`
}

func InitConf() (c Configure) {
	var config Configure
	filename, _ := filepath.Abs("./conf/conf.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	return config
}
