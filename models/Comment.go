package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	PostID  uint
	//User    User `gorm:"foreignKey:UserID"`
	//Post    Post `gorm:"foreignKey:PostID"`
	Post Post `gorm:"references:ID"`
	User User `gorm:"references:ID"`
}
