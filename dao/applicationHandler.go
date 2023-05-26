package dao

import (
	"log"
	"oa-review/db"
	v1 "oa-review/models/protoreq/v1"
)

/*
*
App 实体创建一个新的 Application 并且返回 app id
*/
func (*ApplicationDao) CreateApplication(app *v1.Application) (int64, error) {
	result := db.GetDB().Create(&app)
	if result.Error != nil {
		return -1, result.Error
	}
	return app.Id, nil
}

/*
*
根据 appId 返回 Application 实体
*/
func (dao *ApplicationDao) FindApplicationById(appId int64) (*v1.Application, error) {
	app := v1.Application{Id: appId}
	res := db.GetDB().Where("id = ?", appId).First(&app)
	if res.Error != nil {
		log.Printf("Error on find app by app_id: %v\n", res.Error.Error())
		return nil, res.Error
	}
	status, err := dao.UpdateReviewStatusForApplication(appId)
	if err != nil {
		return nil, err
	}
	app.ReviewStatus = status
	return &app, nil
}

/*
为 application 中的 ApprovedReviewer 添加通过的 Reviewer UserId
*/
func (*ApplicationDao) UpdateApprovedReviewerForApplication(appId int64, reviewerId int64, reviewStatus bool) error {
	app := v1.Application{Id: appId}
	res := db.GetDB().Where("id = ?", appId).First(&app)
	if res.Error != nil {
		log.Printf("Error on update approved revewer for app: %v\n", res.Error.Error())
		return nil
	}

	if reviewStatus {
		app.ApprovedReviewer[reviewerId] = true
	} else {
		delete(app.ApprovedReviewer, reviewerId)
	}
	db.GetDB().Model(&app).Update("approved_reviewer", app.ApprovedReviewer)
	return nil
}

func (*ApplicationDao) UpdateReviewStatusForApplication(appId int64) (bool, error) {
	app := v1.Application{Id: appId}
	res := db.GetDB().Where("id = ?", appId).First(&app)
	if res.Error != nil {
		log.Printf("Error on update reviewer status for app: %v\n", res.Error.Error())
		return false, nil
	}
	ReviewersCount, err := NewReviewerDaoInstance().TableSize()
	if err != nil {
		return false, err
	}
	if len(app.ApprovedReviewer) == int(ReviewersCount) {
		app.ReviewStatus = true
	} else {
		app.ReviewStatus = false
	}
	db.GetDB().Model(&app).Update("review_status", app.ReviewStatus)
	return app.ReviewStatus, nil
}

func (*ApplicationDao) CheckApplicationExist(Id int64) (bool, error) {
	var Application v1.Application
	res := db.GetDB().Where("id = ?", Id).First(&Application)
	if res.Error != nil {
		if res.Error.Error() == "record not found" {
			log.Printf("Error on check Application exist: %v\n", res.Error.Error())
			return false, nil
		}
		return false, res.Error
	}
	return true, nil
}

func (*ApplicationDao) TableSize() (int64, error) {
	var count int64
	if err := db.GetDB().Unscoped().Model(&v1.Application{}).Count(&count).Error; err != nil {
		log.Printf("Error on counting app table size: %v\n", err)
		return 0, err
	}
	return count, nil
}
