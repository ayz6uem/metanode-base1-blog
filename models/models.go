package models

import (
	"base1-blog/config"

	"github.com/sirupsen/logrus"
)

func InitModels() {
	err := config.DB.AutoMigrate(
		&User{},
		&Post{},
		&Comment{},
	)
	if err != nil {
		logrus.Fatal("数据库模型初始化失败", err)
		return
	}
	logrus.Info("数据库模型初始化成功")
}
