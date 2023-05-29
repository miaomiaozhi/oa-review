package runner

import (
	"flag"
	conf "oa-review/conf"
	"oa-review/logger"
	"os"
)

func MustReadConfigFromCmdFlags() (*conf.OaReviewConf, error) {
	var configFilePath string

	// 将命令行参数绑定到configFilePath变量上
	flag.StringVar(&configFilePath, "config", "", "path to config file")
	flag.Parse()

	// 如果未指定config参数，则输出提示信息
	if configFilePath == "" {
		logger.Fatal("config file path is empty")
	}
	if pathChecker(configFilePath) {
		return conf.Read(configFilePath)
	} else {
		logger.Fatal("illegal config file path")
	}
	// unreachable
	return nil, nil
}

func pathChecker(path string) bool {
	logger.Info("path checker", path)
	// 检查文件是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.Fatalf("File %s does not exist\n", path)
		return false
	}

	// 检查文件权限
	if _, err := os.OpenFile(path, os.O_RDONLY, 0666); err != nil {
		logger.Fatalf("File %s is not readable\n", path)
		return false
	}

	// 检查文件类型
	fileInfo, err := os.Stat(path)
	if err != nil {
		logger.Fatalf("Failed to get file info for %s: %v\n", path, err)
		return false
	}

	if fileInfo.IsDir() {
		logger.Fatalf("%s is not a file\n", path)
		return false
	}

	logger.Info("config path is legal", path)
	return true
}
