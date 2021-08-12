package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	CommentID         uint `gorm:"primary_key; AUTO_INCREMENT"`
	BlogID            uint
	UserID            uint
	CommentContent    string `gorm:"size:128; not null" json:"commentcontent"`
	CommentCreateTime time.Time
	CommentUpdateTime time.Time
}
