package runner

import (
	"flag"
	"oa-review/logger"
)

func ParseCmdArgs() string {
	// 定义命令行参数
	var configFile string
	flag.StringVar(&configFile, "config", "", "配置文件路径")

	// 解析命令行参数
	flag.Parse()

	// 输出配置文件路径
	//fmt.Println("配置文件路径:", configFile)
	logger.Info("config file path is", configFile)
	return configFile
}
