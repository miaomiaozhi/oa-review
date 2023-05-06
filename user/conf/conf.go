package conf

import (
	"log"

	"gopkg.in/ini.v1"
)

const (
	configFilePath = "/home/mozezhao/oa-review/user/conf/config.ini"
)

var (
	Username string
	Password string
	Host     string
	Port     string
	Dbname   string
)

func loadConfigData(iniFile *ini.File) {
	Username = iniFile.Section("mysql").Key("username").String()
	Password = iniFile.Section("mysql").Key("password").String()
	Host = iniFile.Section("mysql").Key("host").String()
	Port = iniFile.Section("mysql").Key("port").String()
	Dbname = iniFile.Section("mysql").Key("dbname").String()
}

func init() {
	// read config file
	iniFile, err := ini.Load(configFilePath)
	if err != nil {
		log.Printf("Error on read config file: %v\n", err)
		return
	}

	loadConfigData(iniFile)
}
