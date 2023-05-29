package bean

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// Application info
type ApproverMap map[int64]bool

func (t *ApproverMap) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}
func (t ApproverMap) Value() (driver.Value, error) {
	return json.Marshal(t)
}

type Application struct {
	Id               int64       `gorm:"primary_key"`
	Context          string      `gorm:"default:(-)"`
	ReviewStatus     bool        `gorm:"default:(-)"`
	UserId           int64       `gorm:"default:(-)"`
	ApprovedReviewer ApproverMap `gorm:"default:(-)"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

func (Application) TableName() string {
	return "tb_application"
}
