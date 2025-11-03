package main

import (
	"base1-blog/ui"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Use(web.GlobalRecover())

	ui.LoginController(router)
	ui.PostController(router)

	err := router.Run()
	if err != nil {
		return
	}
}
