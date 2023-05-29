package v1

import (
	"fmt"
	conf "oa-review/conf"
	"oa-review/logger"

	"github.com/tidwall/gjson"
)

const (
	workFlowNamePath   = "workflow.name"
	workFlowStagesPath = "workflow.stages"
)

type PassConditionType int

// 通过条件
const (
	ALL PassConditionType = iota // 全部通过
	HAS                          // 至少一个通过
)

type Reviewer struct {
	Id     int64 // 审核人 id
	Status bool  // 审核状态
}

type Stage struct {
	Reviewers     []*Reviewer       // 审核人群
	Status        bool              // 当前阶段是否通过
	PassCondition PassConditionType // 通过条件
}

type WorkFlow struct {
	name   string   // 流程名称
	index  int32    // 当前流程下标
	stages []*Stage // 所有流程情况
	status bool     // 流程是否终止
}

var workflow *WorkFlow

func getStage(conf *conf.OaReviewConf) []*Stage {
	res := conf.MustGetAny(workFlowStagesPath)
	stages := make([]*Stage, 0)
	res.ForEach(func(k, v gjson.Result) bool {
		reviewersJson := v.Get("reviewers").Array()
		stage := &Stage{
			Reviewers:     make([]*Reviewer, 0),
			Status:        false,
			PassCondition: PassConditionType(v.Get("condition").Num),
		}
		for _, v := range reviewersJson {
			stage.Reviewers = append(stage.Reviewers, &Reviewer{
				Id:     v.Get("id").Int(),
				Status: false,
			})
		}
		stages = append(stages, stage)
		return true
	})
	return stages
}

func (w *WorkFlow) Print() {
	logger.Debug(w.name)
	logger.Debug(w.index)
	for _, v := range w.stages {
		logger.Debug("reviewer cond", v.PassCondition)
		for _, r := range v.Reviewers {
			logger.Debug("reviewer info", r.Id, r.Status)
		}
		fmt.Println()
	}
}

// 初始化工作流程
func InitWorkFlow(conf *conf.OaReviewConf) {
	stages := getStage(conf)
	if len(stages) == 0 {
		workflow = &WorkFlow{
			name:   conf.GetString(workFlowNamePath, ""),
			index:  0,
			stages: stages,
		}
	} else {
		workflow = &WorkFlow{
			name:   conf.GetString(workFlowNamePath, ""),
			index:  0,
			stages: stages,
		}
	}
}

func GetWorkFlow() (bool, *WorkFlow) {
	if workflow == nil {
		logger.Fatal("init workflow must called before get")
		return false, nil
	}
	return workflow.status, workflow
}

func (w *WorkFlow) GetCurentStage() *Stage {
	if w == nil {
		logger.Error("get stage must call before get stage")
		return nil
	}
	if w.status {
		logger.Info("workflow finish")
		return nil
	}
	if len(w.stages) == 0 {
		logger.Info("work flow stages are empty")
		return nil
	}
	return w.stages[w.index]
}

func (w *WorkFlow) SetCurentStage(cur *Stage) bool {
	if w == nil {
		logger.Error("set stage must call before get stage")
		return false
	}
	if cur == nil {
		logger.Error("cur stage is nil, set curent stage error")
		return false
	}
	if len(w.stages) == 0 {
		return true
	}
	w.stages[w.index] = cur
	return w.CheckFinish()
}

func (w *WorkFlow) CheckFinish() bool {
	if w == nil {
		logger.Error("workflow empty, check finish error")
		return false
	}
	for i, v := range w.stages {
		if !v.Status {
			w.status = false
			w.index = int32(i)
			return false
		}
	}
	w.index = int32(len(w.stages))
	w.status = true
	return true
}

func (r *WorkFlow) MoveNext() {
	if r.index < int32(len(r.stages)) {
		r.index += 1
	} else {
		logger.Error("stage finish move next error")
	}
}

func (r *WorkFlow) MovePrev() {
	if r.index > 0 {
		r.index -= 1
	} else {
		logger.Error("stage begin move prev error")
	}
}
