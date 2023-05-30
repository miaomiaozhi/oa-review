package v1

import (
	"fmt"
	conf "oa-review/conf"
	"oa-review/logger"

	"github.com/tidwall/gjson"
)

var (
	workFlowNamePath        = "workflow.name"
	workFlowStagesPath      = "workflow.stages"
	workFlowContextPath     = "workflow.context"
	workFlowApplicantIdPath = "workflow.applicantId"
)

type WorkFlowRunner struct {
}

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
	PassCount     int32             // 通过人数
}

type WorkFlow struct {
	name          string   // 流程名称
	applicantId   int64    // 申请人编号
	applicationId int64    // 申请编号，流程启动时赋值
	context       string   // 申请内容
	index         int32    // 当前流程下标 -1 表示流程未开始，index == len(stages) 表示流程完成
	stages        []*Stage // 所有流程情况
	status        bool     // 流程是否终止

}

var workflow *WorkFlow

// 初始化工作流程
func InitWorkFlow(conf *conf.OaReviewConf) {
	stages := getStage(conf)
	workflow = &WorkFlow{
		name:        conf.GetString(workFlowNamePath, ""),
		applicantId: conf.MustGetInt(workFlowApplicantIdPath),
		context:     conf.MustGetString(workFlowContextPath),
		index:       -1,
		stages:      stages,
	}
	logger.Info("workflow init success")
}

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
		// logger.Info("stage pass condition", stage.PassCondition)
		for _, v := range reviewersJson {
			// logger.Info("reviewer info", v.Get("id").Int())
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

// 返回 workflow, 以及 IsFinish
func GetWorkFlow() (*WorkFlow, bool) {
	if workflow == nil {
		logger.Error("init workflow must called before get")
		return nil, false
	}
	return workflow, workflow.status
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

// 更新当前的状态，并且返回当前是否通过 
func (s *Stage) Pass() bool {
	if s == nil {
		logger.Error("curent stage is empty")
		return false
	}
	if s.PassCondition == ALL {
		s.Status = (s.PassCount == int32(len(s.Reviewers)))
	} else {
		s.Status = (s.PassCount > 0)
	}
	return s.Status
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
