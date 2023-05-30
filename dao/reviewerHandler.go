package dao

import (
	bean "oa-review/bean"
	"oa-review/db"
	"oa-review/logger"
)

/*
*
根据用户名和密码，创建一个新的 Reviewer，返回 Reviewer ID
*/
func (*ReviewerDao) CreateReviewer(reviewer *bean.Reviewer) (int64, error) {
	result := db.GetDB().Create(&reviewer)
	if result.Error != nil {
		return -1, result.Error
	}
	return reviewer.Id, nil
}

/*
*
根据 Id 返回对应的 Reviewer 实体
*/
func (*ReviewerDao) FindReviewerById(reviewerId int64) (*bean.Reviewer, error) {
	reviewer := bean.Reviewer{Id: reviewerId}
	res := db.GetDB().Where("id = ?", reviewerId).First(&reviewer)
	if res.Error != nil {
		logger.Errorf("Error on find reviewer by id: %v\n", res.Error.Error())
		return nil, res.Error
	}
	logger.Info("find reviewer by id", reviewerId)
	return &reviewer, nil
}

/*
传入 Id 以及 option 给 reviewer 的 options 添加操作
*/
func (*ReviewerDao) AddReviewerOption(reviewerId int64, option *bean.ReviewOption) error {
	reviewer := bean.Reviewer{Id: reviewerId}
	res := db.GetDB().Where("id = ?", reviewerId).First(&reviewer)
	if res.Error != nil {
		logger.Errorf("Error on add reviewer option: %s", res.Error.Error())
		return res.Error
	}
	reviewer.Options = append(reviewer.Options, option)
	db.GetDB().Model(&reviewer).Update("options", reviewer.Options)
	return nil
}

func (*ReviewerDao) DeleteReviewerOption(reviewerId int64) (*bean.ReviewOption, error) {
	reviewer := bean.Reviewer{Id: reviewerId}
	res := db.GetDB().Where("id = ?", reviewerId).First(&reviewer)
	if res.Error != nil {
		logger.Errorf("Error on delete reviewer option: %s", res.Error.Error())
		return nil, res.Error
	}
	if len(reviewer.Options) == 0 {
		logger.Error("Error on delete reviewer option: options empty")
		return nil, nil
	}
	optLen := len(reviewer.Options)
	opt := reviewer.Options[optLen-1]
	reviewer.Options = reviewer.Options[:optLen-1]
	db.GetDB().Model(&reviewer).Update("options", reviewer.Options)
	return opt, nil
}

func (*ReviewerDao) CheckReviewerExist(Id int64) (bool, error) {
	var Reviewer bean.Reviewer
	res := db.GetDB().Where("id = ?", Id).Limit(1).Find(&Reviewer)
	if res.Error != nil {
		logger.Errorf("check reviewer exist error: %v", res.Error)
		return false, res.Error
	}
	if Reviewer.Id == 0 {
		return false, nil
	}
	return true, nil
}

func (*ReviewerDao) TableSize() (int64, error) {
	var count int64
	if err := db.GetDB().Unscoped().Model(&bean.Reviewer{}).Count(&count).Error; err != nil {
		logger.Errorf("Error on counting reviewer table size: %v\n", err)
		return 0, err
	}
	return count, nil
}
