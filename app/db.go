package app

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	instance, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic(err)
		return
	}
	instance.AutoMigrate(&User{})
	instance.AutoMigrate(&Comment{})
	instance.AutoMigrate(&Post{})
	db = instance
}
