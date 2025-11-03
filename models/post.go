package models

import (
	"base1-blog/config"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string
	UserID  uint `gorm:"not null"`
}

func PostPage(page, pageSize int) ([]Post, int64) {
	var posts []Post
	var total int64
	config.DB.Model(&Post{}).Count(&total)
	config.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&posts)
	return posts, total
}
