package app

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

func CommentCreate(vo CommentCreateVO) {
	comment := &Comment{Content: vo.Content, UserID: vo.UserID, PostID: vo.PostID}
	db.Create(comment)
}

func CommentList(postId uint) []Comment {
	var comments []Comment
	db.Where("post_id = ?", postId).Find(&comments)
	return comments
}
