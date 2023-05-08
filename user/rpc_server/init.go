package rpcserver

import (
	"log"
	model "oa-review/user/model"
)

// tmp cache
var Users map[int64]*model.User
var Reviewers map[int64]*model.Reviewer
var AppList map[int64]*model.Application
var UserMapping map[int64]int32

func init() {
	log.Println("init user server tmp cache")
	Users = make(map[int64]*model.User)
	Reviewers = make(map[int64]*model.Reviewer)
	AppList = make(map[int64]*model.Application)
	UserMapping = make(map[int64]int32)
}
