package dao

import (
	"errors"
	"log"
	v1 "oa-review/models/protoreq/v1"
)

/*
*
根据用户名和密码，创建一个新的 Reviewer，返回 Reviewer ID
*/
func (*ReviewerDao) CreateReviewer(reviewer *v1.Reviewer) (int64, error) {
	result := DB.Create(&reviewer)
	if result.Error != nil {
		return -1, result.Error
	}
	return reviewer.Id, nil
}

/*
*
根据 Id 返回对应的 Reviewer 实体
*/
func (*ReviewerDao) FindReviewerById(reviewerId int64) (*v1.Reviewer, error) {
	reviewer := v1.Reviewer{Id: reviewerId}
	res := DB.Where("reviewer_id = ?", reviewerId).First(&reviewer)
	if res.Error != nil {
		log.Printf("Error on find reviewer by id: %v\n", res.Error.Error())
		return nil, res.Error
	}
	return &reviewer, nil
}

/*
传入 Id 以及 option 给 reviewer 的 options 添加操作
*/
func (*ReviewerDao) AddReviewerOption(reviewerId int64, option *v1.ReviewOption) error {
	reviewer := v1.Reviewer{Id: reviewerId}
	res := DB.Where("reviewer_id = ?", reviewerId).First(&reviewer)
	if res.Error != nil {
		log.Printf("Error on add reviewer option: %s", res.Error.Error())
		return res.Error
	}
	reviewer.Options = append(reviewer.Options, option)
	DB.Model(&reviewer).Update("options", reviewer.Options)
	return nil
}

func (*ReviewerDao) DeleteReviewerOption(reviewerId int64) (*v1.ReviewOption, error) {
	reviewer := v1.Reviewer{Id: reviewerId}
	res := DB.Where("reviewer_id = ?", reviewerId).First(&reviewer)
	if res.Error != nil {
		log.Printf("Error on delete reviewer option: %s", res.Error.Error())
		return nil, res.Error
	}
	if len(reviewer.Options) == 0 {
		log.Printf("Error on delete reviewer option: options empty")
		return nil, errors.New("Reviewer options empty")
	}
	optLen := len(reviewer.Options)
	opt := reviewer.Options[optLen-1]
	reviewer.Options = reviewer.Options[:optLen-1]
	DB.Model(&reviewer).Update("options", reviewer.Options)
	return opt, nil
}

func (*ReviewerDao) CheckReviewerExist(Id int64) (bool, error) {
	var Reviewer v1.Reviewer
	res := DB.Where("reviewer_id = ?", Id).First(&Reviewer)
	if res.Error != nil {
		if res.Error.Error() == "record not found" {
			log.Printf("Error on check Reviewer exist: %v\n", res.Error.Error())
			return false, nil
		}
		return false, res.Error
	}
	return true, nil
}

func (*ReviewerDao) TableSize() (int64, error) {
	var count int64
	if err := DB.Unscoped().Model(&v1.Reviewer{}).Count(&count).Error; err != nil {
		log.Printf("Error on counting reviewer table size: %v\n", err)
		return 0, err
	}
	return count, nil
}
