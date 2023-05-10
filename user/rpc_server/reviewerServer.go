package rpcserver

import (
	"context"
	"fmt"
	dao "oa-review/dao"
	middleware "oa-review/middleware"
	services "oa-review/proto/services"
)

/*
提交申请
*/
func (userService *UserService) SubmitReview(ctx context.Context, req *services.UserSubmitReviewRequest) (*services.UserSubmitReviewResponse, error) {
	// 用户鉴权
	tokenId, tokenPri, err := middleware.ReviewerAuthorize(req.Token)
	if err != nil {
		return &services.UserSubmitReviewResponse{
			StatusCode: 400,
			StatusMsg:  fmt.Sprintf("Error on submit review: %v", err.Error()),
		}, nil
	}
	if tokenId != req.UserId || tokenPri <= 0 {
		return &services.UserSubmitReviewResponse{
			StatusCode: 400,
			StatusMsg:  "no authorization submit review",
		}, nil
	}

	application, _ := dao.NewApplicationDaoInstance().FindApplicationById(req.ApplicationId)
	_, already := application.ApprovedReviewer[req.UserId]
	if req.ReviewStatus && !already || !req.ReviewStatus && already {
		dao.NewReviewerDaoInstance().AddReviewerOption(req.UserId, &dao.ReviewOption{
			ApplicationId: application.ApplicationId,
			ReviewStatus:  req.ReviewStatus,
		})
	}
	if req.ReviewStatus {
		dao.NewApplicationDaoInstance().UpdateApprovedReviewerForApplication(req.ApplicationId, req.UserId, true)
	} else {
		dao.NewApplicationDaoInstance().UpdateApprovedReviewerForApplication(req.ApplicationId, req.UserId, false)
	}
	dao.NewApplicationDaoInstance().UpdateReviewStatusForApplication(req.ApplicationId)
	return &services.UserSubmitReviewResponse{
		StatusCode: 200,
		StatusMsg:  "ok",
	}, nil
}

func (*UserService) WithdrawReview(ctx context.Context, req *services.UserWithdrawReviewRequest) (*services.UserWithdrawReviewResponse, error) {
	ErrResponse := func(errorMsg string) (*services.UserWithdrawReviewResponse, error) {
		return &services.UserWithdrawReviewResponse{
			StatusCode: 400,
			StatusMsg:  errorMsg,
		}, nil
	}

	// 用户鉴权
	tokenId, tokenPri, err := middleware.ReviewerAuthorize(req.Token)
	if err != nil {
		return &services.UserWithdrawReviewResponse{
			StatusCode: 400,
			StatusMsg:  fmt.Sprintf("Error on submit review: %v", err.Error()),
		}, nil
	}
	if tokenId != req.UserId || tokenPri <= 0 {
		return &services.UserWithdrawReviewResponse{
			StatusCode: 400,
			StatusMsg:  "no authorization submit review",
		}, nil
	}

	option, err := dao.NewReviewerDaoInstance().DeleteReviewerOption(req.UserId)
	if err != nil {
		return ErrResponse(err.Error())
	}
	appId, sta := option.ApplicationId, option.ReviewStatus
	if sta {
		dao.NewApplicationDaoInstance().UpdateApprovedReviewerForApplication(appId, req.UserId, false)
	} else {
		dao.NewApplicationDaoInstance().UpdateApprovedReviewerForApplication(appId, req.UserId, true)
	}
	dao.NewApplicationDaoInstance().UpdateReviewStatusForApplication(appId)

	return &services.UserWithdrawReviewResponse{
		StatusCode: 200,
		StatusMsg:  "withdraw review successfully",
	}, nil
}
