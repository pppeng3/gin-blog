package model

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	UserID         uint
	Title          string `gorm:"size:128; not null" json:"title"`
	Content        string `gorm:"type:longtext; not null" json:"content"`
	BlogCreateTime time.Time
	BlogUpdateTime time.Time
	Like           int32 `gorm:"" json:"like"`
	Dislike        int32 `gorm:"" json:"unlike"`
	Read           int32 `gorm:"" json:"read"`
}
