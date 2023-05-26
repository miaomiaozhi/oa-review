package runner

import (
	conf "oa-review/conf"
	"oa-review/db"
)

type Runner struct {
	Conf *conf.OaReviewConf
}

// TODO add cmd flags
const filePath = "./conf/config.json"

func (r *Runner) Run() {
	// runner.Runner{}.Run()
	confInfo, err := conf.Read(filePath)
	if err != nil {
		panic(err)
	}

	// 初始化配置文件
	conf.InitGlobalConfig(confInfo)
	// 初始化数据库
	db.InitDataBase(conf.GetConfig().Conf)
}
