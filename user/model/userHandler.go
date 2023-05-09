package model

import (
	"log"
	"time"

	"gorm.io/gorm"
)

// User info
type User struct {
	UserId       int64   `gorm:"primary_key"`
	Password     string  `gorm:"default:(-)"`
	Name         string  `gorm:"default:(-)"`
	Applications []int64 `gorm:"default:(-)"`
	Priority     int32   `gorm:"default:(-)"`
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
根据用户名，查找用户实体
*/
func (*UserDao) FindUserByName(username string) (*User, error) {
	user := User{Name: username}

	result := DB.Where("name = ?", username).First(&user)
	err := result.Error
	if err != nil {
		log.Printf("Error on find user by name:%v\n", err)
		return nil, err
	}
	return &user, nil
}

/*
*
根据用户id，查找用户实体
*/
func (*UserDao) FindUserById(id int64) (*User, error) {
	user := User{UserId: id}

	result := DB.Where("user_id = ?", id).First(&user)
	err := result.Error
	if err != nil {
		log.Printf("Error on find user by id:%v\n", err)
		return nil, err
	}
	return &user, nil
}
