package middleware

import (
	"base1-blog/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 检查后续接口是否已登录
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.RequestToken(c)
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "未认证"})
			c.Abort()
			return
		}
		id, err := strconv.Atoi(claims.Subject)
		c.Set("OperatorId", uint(id))
		c.Next()
	}
}

// CurrentUserId 获取当前登录用户ID
func CurrentUserId(c *gin.Context) uint {
	return c.GetUint("OperatorId")
}
