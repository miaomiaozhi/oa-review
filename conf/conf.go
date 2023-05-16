package conf

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	// oa-review
	configFilePath = "/conf/config.yaml"
)

type Config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
}

type ConfigData struct {
	MysqlData Config `yaml:"mysql"`
}

var (
	Username string
	Password string
	Host     string
	Port     string
	Dbname   string
)

var configData ConfigData

func loadConfigData() {
	config := configData.MysqlData

	// read data
	Username = config.Username
	Password = config.Password
	Host = config.Host
	Port = config.Port
	Dbname = config.Dbname
	log.Println("loading config data successfully")
	log.Println("Config data", Username, Password, Host, Port, Dbname)
}

func init() {
	log.Println("Reading database config file")

	dirPath, err := os.Getwd()
	log.Println("dirPath is", dirPath)
	if err != nil {
		log.Printf("Error on reading yaml file: %v\n", err)
		return
	}
	// yamlFile, err := ioutil.ReadFile(configFilePath)
	yamlFileHandle, err := os.Open(dirPath + configFilePath)
	if err != nil {
		log.Printf("Error on reading yaml file: %v\n", err)
		return
	}
	defer yamlFileHandle.Close()

	// 读取配置文件内容
	yamlFile, err := ioutil.ReadAll(yamlFileHandle)
	if err != nil {
		log.Printf("Error on reading yaml file: %v\n", err)
		return
	}
	if err != nil {
		log.Printf("Error on reading yaml file: %v\n", err)
		return
	}
	configData = ConfigData{MysqlData: Config{}}

	err = yaml.Unmarshal(yamlFile, &configData)
	if err != nil {
		log.Printf("Error on unmrashal yaml file: %v\n", err)
		return
	}
	loadConfigData()
}
