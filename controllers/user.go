package controllers

import (
	"base1-blog/middleware"
	"base1-blog/models"
	"base1-blog/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (controller *UserController) Info(c *gin.Context) {
	id := middleware.CurrentUserId(c)
	user, err := models.GetUserById(id)
	if err != nil {
		utils.BadRequest(c, "用户不存在")
		return
	}
	utils.Data(c, user)
}
