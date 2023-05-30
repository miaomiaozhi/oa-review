package v1

import (
	"oa-review/bean"
	"oa-review/dao"
	"oa-review/internal/wrapper"
	"oa-review/logger"
	v1_req "oa-review/models/protoreq/v1"
)

func ReviewerSubmit(ctx *wrapper.Context, reqBody interface{}) error {
	logger.Info("handle ReviewerSubmit now")
	req := reqBody.(*v1_req.ReviewerSubmitRequest)
	// validator

	workflow, finish := GetWorkFlow()
	if workflow == nil {
		wrapper.SendApiBadRequestResponse(ctx, nil, "审核流程未开始")
		return nil
	}
	if finish {
		wrapper.SendApiBadRequestResponse(ctx, nil, "审核流程已结束")
		return nil
	}
	curStage := workflow.GetCurentIndex()
	logger.Debug(curStage)
	workflow.Print()
	err := workflow.SubmitReview(req.UserId, req.ApplicationId, req.ReviewStatus)
	workflow.Print()

	if err != nil {
		wrapper.SendApiBadRequestResponse(ctx, nil, err.Error())
		return nil
	}

	application, _ := dao.NewApplicationDaoInstance().FindApplicationById(req.ApplicationId)
	_, already := application.ApprovedReviewer[req.UserId]
	if req.ReviewStatus && !already || !req.ReviewStatus && already {
		dao.NewReviewerDaoInstance().AddReviewerOption(req.UserId, &bean.ReviewOption{
			Stage:         curStage,
			ApplicationId: application.Id,
			ReviewStatus:  req.ReviewStatus,
		})
	}
	if req.ReviewStatus {
		dao.NewApplicationDaoInstance().UpdateApprovedReviewerForApplication(req.ApplicationId, req.UserId, true)
	} else {
		dao.NewApplicationDaoInstance().UpdateApprovedReviewerForApplication(req.ApplicationId, req.UserId, false)
	}
	dao.NewApplicationDaoInstance().UpdateReviewStatusForApplication(req.ApplicationId)

	wrapper.SendApiOKResponse(ctx, nil, "审核提交成功")
	logger.Info("Apiwrapper reviewer submit review ok")
	return nil
}

func ReviewerWithDraw(ctx *wrapper.Context, reqBody interface{}) error {
	logger.Info("handle ReviewerWithDraw now")

	// TODO: validator
	req := reqBody.(*v1_req.ReviewerWithDrawRequest)

	// 持久化
	option, err := dao.NewReviewerDaoInstance().DeleteReviewerOption(req.UserId)
	if err != nil {
		return err
	}
	if option == nil {
		wrapper.SendApiBadRequestResponse(ctx, nil, "未审核，无法回滚")
		return nil
	}

	// 上一步操作的 appid，以及上一步最终状态 sta
	appId, sta, stage := option.ApplicationId, option.ReviewStatus, option.Stage
	workflow, _ := GetWorkFlow()
	if sta {
		// 将通过修改成不通过
		workflow.WithDrawReview(appId, sta, stage)
		dao.NewApplicationDaoInstance().UpdateApprovedReviewerForApplication(appId, req.UserId, false)
	} else {
		// 将通过修改成不通过
		workflow.WithDrawReview(appId, sta, stage)
		dao.NewApplicationDaoInstance().UpdateApprovedReviewerForApplication(appId, req.UserId, true)
	}
	dao.NewApplicationDaoInstance().UpdateReviewStatusForApplication(appId)

	wrapper.SendApiOKResponse(ctx, nil, "审核回滚成功")
	logger.Info("Apiwrapper reviewer withdraw review ok")
	return nil
}
