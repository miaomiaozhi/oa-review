package runner

import (
	conf "oa-review/conf"
	web "oa-review/web"
)

type Runner struct {
	Conf *conf.OaReviewConf
}

func New(conf *conf.OaReviewConf) *Runner {
	return &Runner{
		Conf: conf,
	}
}

// 启动 app
func (r *Runner) StartWebApp() {
	web.New().Run(r.Conf)
}
