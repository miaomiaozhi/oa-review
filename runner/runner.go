package runner

import (
	conf "oa-review/conf"
	"oa-review/db"
	workflow "oa-review/services/v1"
)

type Runner struct {
	conf *conf.OaReviewConf
}

// TODO add cmd flags
const filePath = "./conf/config.json"

func (r *Runner) Run() {
	confInfo, err := conf.Read(filePath)
	if err != nil {
		panic(err)
	}
	r.conf = confInfo

	// 初始化配置文件
	conf.InitGlobalConfig(confInfo)
	// 初始化数据库
	db.InitDataBase(conf.GetConfig().Conf)
	// 初始化工作流
	workflow.InitWorkFlow(conf.GetConfig().Conf)
	_, w := workflow.GetWorkFlow()
	w.Print()
}
