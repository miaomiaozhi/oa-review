package v1

import (
	"fmt"
	"oa-review/logger"
)

func (w *WorkFlow) CanStart() bool {
	return w.index == -1
}

func (w *WorkFlow) IsStarted() bool {
	return 0 <= w.index && w.index < int32(len(w.stages))
}

func (w *WorkFlow) Start() bool {
	if w.CanStart() {
		w.index = 0
		return true
	}
	return false
}

func (w *WorkFlow) GetCurentIndex() int32 {
	return w.index
}

func (w *WorkFlow) GetApplicantId() int64 {
	return w.applicantId
}

func (w *WorkFlow) GetContext() string {
	return w.context
}

func (w *WorkFlow) CheckWorkFlowExist(userId int64, ctx string) bool {
	return w.applicantId == userId && w.context == ctx
}

func (w *WorkFlow) SetApplicationId(appId int64) {
	if w == nil {
		logger.Error("init workflow must be called before set")
		return
	}
	w.applicationId = appId
}

func (w *WorkFlow) CheckCurrentStageFinish() bool {
	if w == nil {
		logger.Error("workflow is nil")
		return false
	}
	if w.status {
		logger.Info("workflow is finish")
		return true
	}
	if !w.IsStarted() {
		logger.Info("workflow is not started")
		return false
	}
	return w.GetCurentStage().Pass()
}

// 提交审核 返回该操作是否成功
func (w *WorkFlow) SubmitReview(reviewerId int64, applicationId int64, status bool) error {
	if !w.IsStarted() {
		logger.Error("workflow submit review error")
		return fmt.Errorf("流程未开始")
	}
	logger.Info("workflow submit review")
	stage := w.GetCurentStage()
	msg := stage.Update(reviewerId, status)
	if stage.Pass() {
		w.MoveNext()
	}
	return msg
}

// 将当前的阶段更新，
func (s *Stage) Update(reviewerId int64, status bool) error {
	for i, v := range s.Reviewers {
		if v.Id == reviewerId {
			if s.Reviewers[i].Status != status {
				if s.Reviewers[i].Status {
					s.PassCount -= 1
				} else {
					s.PassCount += 1
				}
			}
			s.Reviewers[i].Status = status
			logger.Debug("stage update success", reviewerId, status)
			_ = s.Pass()
			return nil
		}
	}
	logger.Debug("stage update error", reviewerId, status)
	return fmt.Errorf("无法审核当前阶段")
}

func (w *WorkFlow) updateAllStages() {
	if !w.IsStarted() {
		logger.Error("workflow update all stages error: workflow is not started")
	}
	idx := int32(0)
	for int(idx) < len(w.stages) {
		if w.stages[idx].Pass() {
			idx += 1
		} else {
			break
		}
	}
	for int(idx) < len(w.stages) {
		w.stages[idx].Status = false
	}
	w.index = idx
}

func (w *WorkFlow) WithDrawReview(reviewerId int64, status bool, stageIdx int32) bool {
	if !w.IsStarted() {
		logger.Error("workflow with draw review error: is not started")
		return false
	}
	logger.Info("workflow with draw review")
	if stageIdx < 0 || stageIdx >= int32(len(w.stages)) {
		logger.Error("workflow with draw review error: stage index illegal")
		return false
	}
	stage := w.stages[stageIdx]
	stage.Update(reviewerId, status)
	_ = stage.Pass()
	w.updateAllStages()
	return true
}
