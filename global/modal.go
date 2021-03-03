package global

import (
	"gorm.io/gorm"
)

type GMODEL struct {
	ID          uint           `gorm:"primarykey"`
	CreateTime  MyTime         `json:"create_time"`
	UpdateTime  MyTime         `json:"update_time"`
	DeletedTime gorm.DeletedAt `gorm:"index" json:"-"`
}
