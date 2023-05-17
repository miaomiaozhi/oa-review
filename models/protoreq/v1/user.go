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
