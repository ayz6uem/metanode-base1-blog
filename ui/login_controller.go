package ui

import (
	"base1-blog/app"

	"github.com/gin-gonic/gin"
)

func LoginController(router *gin.Engine) {
	router.POST("/register", register)
	router.POST("/login", login)
	router.GET("/my", app.Logined(), info)
}

func register(c *gin.Context) {
	vo := app.UserRegister{}
	err := c.ShouldBind(&vo)
	if err != nil {
		panic(err)
		return
	}
	app.Register(vo)
	c.JSON(200, gin.H{"code": 0, "msg": "ok"})
}

func login(c *gin.Context) {
	var vo app.UserLogin
	err := c.ShouldBindJSON(&vo)
	if err != nil {
		panic(err)
		return
	}
	token := app.Login(vo)
	c.JSON(200, gin.H{"code": 0, "msg": "ok", "token": token})
}

func info(c *gin.Context) {
	id := c.GetUint("OperatorId")
	user := app.Get(id)
	c.JSON(200, gin.H{"code": 0, "msg": "ok", "user": user})
}
