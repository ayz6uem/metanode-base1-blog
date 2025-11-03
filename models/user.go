package models

import (
	"base1-blog/config"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
	Email    string
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// UserExists 根据用户名查询用户是否存在
func UserExists(Username string) bool {
	var count int64
	config.DB.Model(&User{}).Where("username = ?", Username).Count(&count)
	return count > 0
}

// GetUserByUsername 根据用户名查询用户，仅在登录时使用
func GetUserByUsername(Username string) *User {
	user := &User{}
	result := config.DB.Model(&User{}).Where("username = ?", Username).First(user)
	if result.Error != nil {
		panic("未找到用户")
	}
	return user
}

// GetUserById 根据用户ID查询用户
func GetUserById(id uint) (*User, error) {
	user := &User{}
	result := config.DB.Model(&User{}).Where("id = ?", id).First(user)
	if result.Error != nil {
		return nil, errors.New("未找到用户")
	}
	return user, nil
}
