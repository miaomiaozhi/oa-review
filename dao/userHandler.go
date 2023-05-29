package dao

import (
	bean "oa-review/bean"
	"oa-review/db"
	"oa-review/logger"
)

/*
*
根据用户名和密码，创建一个新的 User，返回UserId
*/
func (*UserDao) CreateUser(user *bean.User) (int64, error) {
	result := db.GetDB().Create(&user)
	logger.Info("user", user)
	if result.Error != nil {
		return -1, result.Error
	}
	return user.Id, nil
}

/*
*
根据用户id，查找用户实体
*/
func (*UserDao) FindUserByUserId(id int64) (*bean.User, error) {
	user := bean.User{Id: id}

	result := db.GetDB().Where("id = ?", id).First(&user)
	err := result.Error
	if err != nil {
		logger.Error("find user by user id error", result.Error.Error())
		return nil, err
	}
	return &user, nil
}

/*
往 user 中的 请求列表中添加一个 appId
*/
func (*UserDao) AddApplicationForUser(userId int64, appId int64) error {
	var user bean.User
	res := db.GetDB().Where("id = ?", userId).First(&user)
	if res.Error != nil {
		logger.Errorf("Error on add app for user: %v\n", res.Error.Error())
		return res.Error
	}
	// logger.Debug("user before appliacitons", user.Applications)

	user.Applications = append(user.Applications, appId)
	// logger.Debug("user after appliacitons", user.Applications)
	db.GetDB().Model(&user).UpdateColumn("applications", user.Applications)
	// db.GetDB().Model(&user).Updates(map[string]interface{}{"applications": user.Applications})
	return nil
}

func (*UserDao) CheckUserExist(userId int64) (bool, error) {
	var user bean.User
	res := db.GetDB().Where("id = ?", userId).Limit(1).Find(&user)
	// logger.Debug("user id", userId)
	// logger.Debug(user)
	if res.Error != nil {
		logger.Error("check user exist error", res.Error.Error())
		return false, res.Error
	}
	// 返回 0 值，查找不到
	if user.Id == 0 {
		return false, nil
	}
	return true, nil
}

func (*UserDao) TableSize() (int64, error) {
	var count int64
	if err := db.GetDB().Unscoped().Model(&bean.User{}).Count(&count).Error; err != nil {
		logger.Errorf("Error on counting user table size: %v\n", err)
		return 0, err
	}
	return count, nil
}
