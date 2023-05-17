package dao

import (
	"log"
	v1 "oa-review/models/protoreq/v1"
)

/*
*
根据用户名和密码，创建一个新的 User，返回UserId
*/
func (*UserDao) CreateUser(user *v1.User) (int64, error) {
	result := DB.Create(&user)
	if result.Error != nil {
		return -1, result.Error
	}
	return user.Id, nil
}

/*
*
根据用户id，查找用户实体
*/
func (*UserDao) FindUserByUserId(id int64) (*v1.User, error) {
	user := v1.User{Id: id}

	result := DB.Where("user_id = ?", id).First(&user)
	err := result.Error
	if err != nil {
		log.Printf("Error on find user by id: %v\n", err)
		return nil, err
	}
	return &user, nil
}

/*
往 user 中的 请求列表中添加一个 appId
*/
func (*UserDao) AddApplicationForUser(userId int64, appId int64) error {
	var user v1.User
	res := DB.Where("user_id = ?", userId).First(&user)
	if res.Error != nil {
		log.Printf("Error on add app for user: %v\n", res.Error.Error())
		return res.Error
	}
	user.Applications = append(user.Applications, appId)
	DB.Model(&user).UpdateColumn("applications", user.Applications)
	return nil
}

func (*UserDao) CheckUserExist(userId int64) (bool, error) {
	var user v1.User
	res := DB.Where("user_id = ?", userId).First(&user)
	if res.Error != nil {
		if res.Error.Error() == "record not found" {
			log.Printf("Error on check user exist: %v\n", res.Error.Error())
			return false, nil
		}
		return false, res.Error
	}
	return true, nil
}

func (*UserDao) TableSize() (int64, error) {
	var count int64
	if err := DB.Unscoped().Model(&v1.User{}).Count(&count).Error; err != nil {
		log.Printf("Error on counting user table size: %v\n", err)
		return 0, err
	}
	return count, nil
}
