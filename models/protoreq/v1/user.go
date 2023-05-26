package v1

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Apps []int64

func (a *Apps) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Apps")
	}
	var ints []int64
	if err := json.Unmarshal(bytes, &ints); err != nil {
		return fmt.Errorf("failed to unmarshal Apps")
	}
	*a = ints
	return nil
}
func (a Apps) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// User info
type User struct {
	Id           int64  `gorm:"primary_key"`
	Password     string `gorm:"default:(-)"`
	Name         string `gorm:"default:(-)"`
	Applications Apps   `gorm:"type:json;default:(-)"`
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
