package bean

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// 上一个审核操作的结构
type ReviewOption struct {
	Stage         int32 // 阶段
	ApplicationId int64 // 申请ID
	ReviewStatus  bool  // 审核状态
}

type ReviewOptions []*ReviewOption

func (t *ReviewOptions) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}
func (t ReviewOptions) Value() (driver.Value, error) {
	return json.Marshal(t)
}

// Reviewer info
type Reviewer struct {
	Id           int64         `gorm:"primary_key"`
	Name         string        `gorm:"default:(-)"`
	Applications Apps          `gorm:"type:json;default:(-)"`
	Options      ReviewOptions `gorm:"type:json;default:(-)"`
	Priority     int32         `gorm:"default:(-)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (Reviewer) TableName() string {
	return "tb_reviewer"
}
