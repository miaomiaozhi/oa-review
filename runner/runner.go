package runner

import "oa-review/web"

type Runner struct {
}

func (Runner) Run() {
	web.OaReviewWeb{}.Run()
}
