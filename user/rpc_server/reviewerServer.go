package rpcserver

import (
	"context"
	dao "oa-review/dao"
	services "oa-review/proto/services"
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
		return ErrResponse("reviewer id illegal")
	}
	// // DAO find usr by id
	// if _, userExist := Users[req.UserId]; !userExist {
	// 	return ErrResponse("user not find")
	// }
	exist, err := dao.NewReviewerDaoInstance().CheckReviewerExist(req.UserId)
	if err != nil {
		return ErrResponse(err.Error())
	}
	if !exist {
		return ErrResponse("reviewer not found")
	}
	exist, err = dao.NewApplicationDaoInstance().CheckApplicationExist(req.ApplicationId)
	if err != nil {
		return ErrResponse(err.Error())
	}
	if !exist {
		return ErrResponse("app not found")
	}

	application, _ := dao.NewApplicationDaoInstance().FindApplicationById(req.ApplicationId)

	// // // 如果是一个有效状态
	// // DAO find app exist
	// _, already := AppList[req.ApplicationId].ApprovedReviewer[req.UserId]
	// if req.ReviewStatus && !already || !req.ReviewStatus && already {
	// 	// DAO updata reviewer
	// 	Reviewers[req.UserId].Options = append(Reviewers[req.UserId].Options, &dao.ReviewOption{
	// 		ApplicationId: req.ApplicationId,
	// 		ReviewStatus:  req.ReviewStatus,
	// 	})
	// }
	_, already := application.ApprovedReviewer[req.UserId]
	if req.ReviewStatus && !already || !req.ReviewStatus && already {
		dao.NewReviewerDaoInstance().AddReviewerOption(req.UserId, &dao.ReviewOption{
			ApplicationId: application.ApplicationId,
			ReviewStatus:  req.ReviewStatus,
		})
	}

	// updata sql data app
	if req.ReviewStatus {
		// DAO update app
		// AppList[req.ApplicationId].ApprovedReviewer[req.UserId] = true
		dao.NewApplicationDaoInstance().UpdateApprovedReviewerForApplication(req.ApplicationId, req.UserId, true)
	} else {
		// DAO update app
		// delete(AppList[req.ApplicationId].ApprovedReviewer, req.UserId)
		dao.NewApplicationDaoInstance().UpdateApprovedReviewerForApplication(req.ApplicationId, req.UserId, false)
	}
	// DAO update app
	// AppList[req.ApplicationId].ReviewStatus = (len(AppList[req.ApplicationId].ApprovedReviewer) == len(Reviewers))
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
	if req.UserId < 0 {
		return ErrResponse("user id illegal")
	}

	exist, err := dao.NewReviewerDaoInstance().CheckReviewerExist(req.UserId)
	if err != nil {
		return ErrResponse(err.Error())
	}
	if !exist {
		return ErrResponse("reviewer not found")
	}

	option, err := dao.NewReviewerDaoInstance().DeleteReviewerOption(req.UserId)
	if err != nil {
		return ErrResponse(err.Error())
	}
	// if len(Reviewers[req.UserId].Options) == 0 {
	// 	return ErrResponse("error on withdraw review, empty options")
	// }

	// optLen := len(Reviewers[req.UserId].Options)
	// option := Reviewers[req.UserId].Options[optLen-1]

	// // 删除最后一个操作
	// Reviewers[req.UserId].Options = Reviewers[req.UserId].Options[:optLen-1]

	// DAO updata reviewer app
	appId, sta := option.ApplicationId, option.ReviewStatus
	if sta {
		// delete(AppList[appId].ApprovedReviewer, req.UserId)
		dao.NewApplicationDaoInstance().UpdateApprovedReviewerForApplication(appId, req.UserId, false)
	} else {
		dao.NewApplicationDaoInstance().UpdateApprovedReviewerForApplication(appId, req.UserId, true)
		// AppList[appId].ApprovedReviewer[req.UserId] = true
	}
	dao.NewApplicationDaoInstance().UpdateReviewStatusForApplication(appId)
	// AppList[appId].ReviewStatus = (len(AppList[appId].ApprovedReviewer) == len(Reviewers))

	return &services.UserWithdrawReviewResponse{
		StatusCode: 200,
		StatusMsg:  "withdraw review successfully",
	}, nil
}
