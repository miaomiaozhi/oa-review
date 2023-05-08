package model

import (
	"log"
	"sync"
	"time"

	"gorm.io/gorm"
)

type UserDao struct{}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

// Application info
type Application struct {
	gorm.Model
	ApplicationId    int64          `gorm:"primary_key"`
	Context          string         `gorm:"default:(-)"`
	ReviewStatus     bool           `gorm:"default:(-)"`
	UserId           int64          `gorm:"default:(-)"`
	ApprovedReviewer map[int64]bool `gorm:"default:(-)"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

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

func NewModelUser() *User {
	return &User{
		Applications: make([]int64, 0),
	}
}

func NewModelApplication() *Application {
	return &Application{
		ApprovedReviewer: make(map[int64]bool, 0),
	}
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
