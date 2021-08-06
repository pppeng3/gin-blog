package model

type BlogUser struct {
	ID           int    `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Username     string `gorm:"column:username;type:char(1);not null" json:"username"`
	Userpassword string `gorm:"column:userpassword;type:char(1);not null" json:"userpassword"`
	Phone        string `gorm:"column:phone;type:char(1)" json:"phone"`
	Email        string `gorm:"column:email;type:char(1)" json:"email"`
}

