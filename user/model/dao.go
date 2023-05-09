package model

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

/*
TODO
mysql handler:

CreateUser(user) (int64, err)
FindUserByUserId(user id) (*User, err)
UpdateUserInfo(user) (err)

CreateReviewer(reviewer) (int64, err)
FindReviewerByUserId(user id) (*Reviewer, err)
UpdataReviewer(reviewer) (err)

CreateApplication(appli) (int64, err)
FindApplicationById(app id) (*App, err)
UpdateApplication(app) (err)
*/
