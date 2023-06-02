package main

import (
	"oa-review/conf"
	"oa-review/db"
	"oa-review/logger"
	"oa-review/runner"
)

func main() {
	// 初始化配置文件
	config, err := runner.MustReadConfigFromCmdFlags()
	if err != nil {
		logger.Fatalf("read config error: %v", err.Error())
	}
	conf.InitGlobalConfig(config)

	// 初始化日志
	logger.InitWithConfig(config)

	// 初始化数据库
	db.InitDataBase(config)

	// 初始化 web
	runner.New(config).StartWebApp()
}
