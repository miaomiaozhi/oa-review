package runner

import (
	"oa-review/conf"
	"oa-review/logger"
)

// 流程信息
type StageConfig struct {
	Reviewer []int64 // 审核人
	Status   []bool
	Finish   bool
}

type ReviewStage struct {
	Applicant int64          // 申请人
	Name      string         // 审核名称
	stages    []*StageConfig // 流程
	cur       *StageConfig   // 当前流程
	index     int32          // 当前的流程
}

func InitReviewStage(conf *conf.OaReviewConf) *ReviewStage {
	return &ReviewStage{
		Name:  conf.MustGetString("stage.name"),
		index: 0,
	}
}

func (r *ReviewStage) GetCurStage() *StageConfig {
	if r == nil || r.cur == nil {
		logger.Error("set stage must call before get stage")
		return nil
	}
	return r.cur
}

func (r *ReviewStage) MoveNext() {
	if r.index < int32(len(r.stages)) {
		r.index += 1
		r.cur = r.stages[r.index]
	} else {
		logger.Error("stage finish move next error")
	}
}

func (r *ReviewStage) MovePrev() {
	if r.index > 0 {
		r.index -= 1
		r.cur = r.stages[r.index]
	} else {
		logger.Error("stage begin move prev error")
	}
}
