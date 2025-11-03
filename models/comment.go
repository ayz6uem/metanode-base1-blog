package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	PostID  uint
}

type CommentCreateVO struct {
	Content string `json:"content"`
	PostID  uint
	UserID  uint
}
