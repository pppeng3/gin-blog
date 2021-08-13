package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	BlogID            uint
	UserID            uint
	CommentContent    string `gorm:"size:128; not null" json:"commentcontent"`
	CommentCreateTime time.Time
	CommentUpdateTime time.Time
}
