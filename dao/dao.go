package dao

import (
	"sync"
)

type UserDao struct{}
type ReviewerDao struct{}
type ApplicationDao struct{}

var userDao *UserDao
var userOnce sync.Once
var reviewerDao *ReviewerDao
var reviewerOnce sync.Once
var applicationDao *ApplicationDao
var applicationOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

func NewReviewerDaoInstance() *ReviewerDao {
	reviewerOnce.Do(func() {
		reviewerDao = &ReviewerDao{}
	})
	return reviewerDao
}

func NewApplicationDaoInstance() *ApplicationDao {
	applicationOnce.Do(func() {
		applicationDao = &ApplicationDao{}
	})
	return applicationDao
}
