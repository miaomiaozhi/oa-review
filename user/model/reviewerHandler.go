package model

import (
	"time"

	"gorm.io/gorm"
)

type ReviewOption struct {
	ApplicationId int64
	ReviewStatus  bool
}

// Reviewer info
type Reviewer struct {
	UserId       int64           `gorm:"primary_key"`
	Name         string          `gorm:"default:(-)"`
	Applications []int64         `gorm:"default:(-)"`
	Options      []*ReviewOption `gorm:"default:(-)"`
	Priority     int32           `gorm:"default:(-)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

/*
根据返回所有审核人用户，列表形式
*/
func (*UserDao) FindReviewerList() ([]*User, error) {
	var reviewer []*User
	result := DB.Where("priority > ?", 0).Order("priority asc").Find(reviewer)
	if result.Error != nil {
		return nil, result.Error
	}
	return reviewer, nil
}
