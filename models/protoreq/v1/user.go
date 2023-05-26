package v1

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Apps []int64

func (t *Apps) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}
func (t *Apps) Value() (driver.Value, error) {
	return json.Marshal(t)
}

// User info
type User struct {
	Id           int64  `gorm:"primary_key"`
	Password     string `gorm:"default:(-)"`
	Name         string `gorm:"default:(-)"`
	Applications Apps   `gorm:"default:(-)"`
	Priority     int32  `gorm:"default:(-)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return "tb_user"
}

// user request
type UserLoginRequest struct {
	UserId       string `json:"UserId,omitempty"`       // user id
	UserPassword string `json:"UserPassword,omitempty"` // user password
}

type UserRegisterRequest struct {
	UserId       string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	UserPassword string `protobuf:"bytes,2,opt,name=UserPassword,proto3" json:"UserPassword,omitempty"`
	UserName     string `protobuf:"bytes,3,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Priority     int32  `protobuf:"varint,4,opt,name=Priority,proto3" json:"Priority,omitempty"`
}
