package dao

import (
	"database/sql/driver"
	"encoding/json"
	"log"
	"time"

	"gorm.io/gorm"
)

type Apps []int64

func (t *Apps) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}
func (t Apps) Value() (driver.Value, error) {
	return json.Marshal(t)
}

// User info
type User struct {
	UserId       int64  `gorm:"primary_key"`
	Password     string `gorm:"default:(-)"`
	Name         string `gorm:"default:(-)"`
	Applications Apps   `gorm:"default:(-)"`
	Priority     int32  `gorm:"default:(-)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

/*
*
根据用户名和密码，创建一个新的 User，返回UserId
*/
func (*UserDao) CreateUser(user *User) (int64, error) {
	result := DB.Create(&user)
	if result.Error != nil {
		return -1, result.Error
	}
	return user.UserId, nil
}

/*
*
根据用户id，查找用户实体
*/
func (*UserDao) FindUserByUserId(id int64) (*User, error) {
	user := User{UserId: id}

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
	var user User
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
	var user User
	res := DB.Where("user_id = ?", userId).First(&user)
	if res.Error != nil {
		if res.Error.Error() != "record not found" {
			log.Printf("Error on check user exist: %v\n", res.Error.Error())
			return false, nil
		}
		return false, res.Error
	}
	return true, nil
}

func (*UserDao) TableSize() (int64, error) {
	var count int64
	if err := DB.Unscoped().Model(&User{}).Count(&count).Error; err != nil {
		log.Printf("Error on counting user table size: %v\n", err)
		return 0, err
	}
	return count, nil
}
