package app

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string
	UserID  uint `gorm:"not null"`
}

type PostCreateVO struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  uint
}

type PostUpdateVO struct {
	ID      uint
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  uint
}

func PostCreate(vo PostCreateVO) {
	post := &Post{Title: vo.Title, Content: vo.Content, UserID: vo.UserID}
	db.Create(post)
}

func PostPage(page, pageSize int) ([]Post, int64) {
	var posts []Post
	var total int64
	db.Model(&Post{}).Count(&total)
	db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&posts)
	return posts, total
}

func PostUpdate(vo PostUpdateVO) {
	post := &Post{}
	result := db.Model(post).Where("id = ?", vo.ID).First(post)
	if result.RowsAffected == 0 {
		panic("文章不存在")
	}
	if post.UserID != vo.UserID {
		panic("没有权限修改此文章")
	}
	post.Title = vo.Title
	post.Content = vo.Content
	db.Save(post)
}

func PostDelete(id uint, operatorId uint) {
	post := &Post{}
	result := db.Model(post).Where("id = ?", id).First(post)
	if result.RowsAffected == 0 {
		panic("文章不存在")
	}
	if post.UserID != operatorId {
		panic("没有权限删除此文章")
	}
	db.Delete(post)
}
