package v1

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type ReviewOption struct {
	ApplicationId int64
	ReviewStatus  bool
}

type ReviewOptions []*ReviewOption

func (t *ReviewOptions) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}
func (t *ReviewOptions) Value() (driver.Value, error) {
	return json.Marshal(t)
}

// Reviewer info
type Reviewer struct {
	Id           int64         `gorm:"primary_key"`
	Name         string        `gorm:"default:(-)"`
	Applications Apps          `gorm:"default:(-)"`
	Options      ReviewOptions `gorm:"default:(-)"`
	Priority     int32         `gorm:"default:(-)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (Reviewer) TableName() string {
	return "tb_reviewer"
}
