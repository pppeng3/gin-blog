package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"" json:"username"`
	Password string `gorm:"" json:"password"`
}

type UserExtra struct {
	gorm.Model
	UserId      int32  `gorm:"" json:"user_id"`
	Nickname    string `gorm:"" json:"nickname"`
	Email       string `gorm:"" json:"email"`
	PhoneNumber string `gorm:"" json:"phone_number"`
	AvatarUrl   string `gorm:"" json:"avatar_url"`
}
