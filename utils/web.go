package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func RequestToken(c *gin.Context) string {
	authorization := c.GetHeader("Authorization")
	if authorization != "" && strings.HasPrefix(authorization, "Bearer ") {
		return strings.TrimPrefix(authorization, "Bearer ")
	}
	return c.Query("token")
}

type PageQuery struct {
	Page     int `form:"page"`
	PageSize int `form:"size"`
}

type EntityIdVO struct {
	ID uint `uri:"id" binding:"required"`
}
