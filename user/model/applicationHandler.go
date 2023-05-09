package model // Application info
import (
	"time"

	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	ApplicationId    int64          `gorm:"primary_key"`
	Context          string         `gorm:"default:(-)"`
	ReviewStatus     bool           `gorm:"default:(-)"`
	UserId           int64          `gorm:"default:(-)"`
	ApprovedReviewer map[int64]bool `gorm:"default:(-)"` // 数据库不支持
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
