package controllers

import (
	"base1-blog/config"
	"base1-blog/models"
	"base1-blog/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
}

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token     string      `json:"token"`
	ExpiresAt time.Time   `json:"expires_at"`
	User      models.User `json:"user"`
}

func (controller *AuthController) Register(c *gin.Context) {
	vo := UserRegisterRequest{}
	if err := c.ShouldBind(&vo); err != nil {
		utils.BadRequest(c, "参数错误")
		return
	}

	if models.UserExists(vo.Username) {
		utils.BadRequest(c, "用户名已存在")
		return
	}
	user := models.User{}
	user.Username = vo.Username
	user.Password = vo.Password
	user.Email = vo.Email
	config.DB.Create(&user)

	token, expiresAt, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.InternalServerError(c, "注册失败")
		return
	}
	utils.Data(c, AuthResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		User:      user,
	})
}

func (controller *AuthController) Login(c *gin.Context) {
	vo := UserLoginRequest{}
	if err := c.ShouldBindJSON(&vo); err != nil {
		utils.BadRequest(c, "参数错误")
		return
	}
	user := models.GetUserByUsername(vo.Username)
	user.CheckPassword(vo.Password)
	token, expiresAt, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.InternalServerError(c, "登录失败")
		return
	}
	utils.Data(c, AuthResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		User:      *user,
	})
}
