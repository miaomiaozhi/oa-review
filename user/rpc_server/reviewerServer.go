package rpcserver

import (
	"context"
	"oa-review/user/model"
	services "oa-review/user/services"
)

/*
提交申请
TODO maintain options
*/
func (userService *UserService) SubmitReview(ctx context.Context, req *services.UserSubmitReviewRequest) (*services.UserSubmitReviewResponse, error) {
	ErrResponse := func(errorMsg string) (*services.UserSubmitReviewResponse, error) {
		return &services.UserSubmitReviewResponse{
			StatusCode: 400,
			StatusMsg:  errorMsg,
		}, nil
	}
	if req.UserId < 0 {
		return ErrResponse("user id illegal")
	}
	if _, userExist := Users[req.UserId]; !userExist {
		return ErrResponse("user not find")
	}
	if _, appExist := AppList[req.ApplicationId]; !appExist {
		return ErrResponse("app not find")
	}

	// 如果是一个有效状态
	_, already := AppList[req.ApplicationId].ApprovedReviewer[req.UserId]
	if req.ReviewStatus && !already || !req.ReviewStatus && already {
		Reviewers[req.UserId].Options = append(Reviewers[req.UserId].Options, &model.ReviewOption{
			ApplicationId: req.ApplicationId,
			ReviewStatus:  req.ReviewStatus,
		})
	}

	// updata sql data app
	if req.ReviewStatus {
		AppList[req.ApplicationId].ApprovedReviewer[req.UserId] = true
	} else {
		delete(AppList[req.ApplicationId].ApprovedReviewer, req.UserId)
	}
	AppList[req.ApplicationId].ReviewStatus = (len(AppList[req.ApplicationId].ApprovedReviewer) == len(Reviewers))

	// update user data applist

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
	if req.UserId < 0 {
		return ErrResponse("user id illegal")
	}
	// DAO
	if _, userExist := Reviewers[req.UserId]; !userExist {
		return ErrResponse("user not find")
	}

	if len(Reviewers[req.UserId].Options) == 0 {
		return ErrResponse("error on withdraw review, empty options")
	}

	optLen := len(Reviewers[req.UserId].Options)
	option := Reviewers[req.UserId].Options[optLen-1]

	// 删除最后一个操作
	Reviewers[req.UserId].Options = Reviewers[req.UserId].Options[:optLen-1]

	appId, sta := option.ApplicationId, option.ReviewStatus
	if sta {
		delete(AppList[appId].ApprovedReviewer, req.UserId)
	} else {
		AppList[appId].ApprovedReviewer[req.UserId] = true
	}
	AppList[appId].ReviewStatus = (len(AppList[appId].ApprovedReviewer) == len(Reviewers))

	return &services.UserWithdrawReviewResponse{
		StatusCode: 200,
		StatusMsg:  "withdraw review successfully",
	}, nil
}
