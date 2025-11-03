package controllers

import (
	"base1-blog/config"
	"base1-blog/middleware"
	"base1-blog/models"
	"base1-blog/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PostController struct {
}

type PostCreateRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  uint
}

type PostUpdateRequest struct {
	ID      uint
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  uint
}

type PostPageResponse struct {
	Total int64         `json:"total"`
	List  []models.Post `json:"list"`
}

func (controller *PostController) Create(c *gin.Context) {
	vo := PostCreateRequest{}
	if err := c.ShouldBind(&vo); err != nil {
		utils.BadRequest(c, "参数错误")
		return
	}
	vo.UserID = middleware.CurrentUserId(c)
	post := &models.Post{Title: vo.Title, Content: vo.Content, UserID: vo.UserID}
	config.DB.Create(post)
	utils.Success(c)
}

func (controller *PostController) Page(c *gin.Context) {
	query := utils.PageQuery{}
	if err := c.ShouldBind(&query); err != nil {
		utils.BadRequest(c, "参数错误")
		return
	}
	content, total := models.PostPage(query.Page, query.PageSize)
	utils.Data(c, PostPageResponse{
		Total: total,
		List:  content,
	})
}

func (controller *PostController) Update(c *gin.Context) {
	entityId := utils.EntityIdVO{}
	vo := PostUpdateRequest{}
	if err := c.ShouldBindUri(&entityId); err != nil {
		utils.BadRequest(c, "参数错误")
		return
	}
	if err := c.ShouldBind(&vo); err != nil {
		utils.BadRequest(c, "参数错误")
		return
	}
	vo.ID = entityId.ID
	vo.UserID = middleware.CurrentUserId(c)

	post := &models.Post{}
	result := config.DB.Model(post).Where("id = ?", vo.ID).First(post)
	if result.Error != nil {
		logrus.Error(result.Error.Error())
		utils.InternalServerError(c, "服务错误")
		return
	}
	if result.RowsAffected == 0 {
		utils.BadRequest(c, "文章不存在")
		return
	}
	if post.UserID != vo.UserID {
		utils.Forbidden(c, "没有权限修改此文章")
		return
	}
	post.Title = vo.Title
	post.Content = vo.Content
	config.DB.Save(post)
	utils.Success(c)
}

func (controller *PostController) Del(c *gin.Context) {
	entityId := utils.EntityIdVO{}
	if err := c.ShouldBindUri(&entityId); err != nil {
		utils.BadRequest(c, "参数错误")
		return
	}

	post := &models.Post{}
	result := config.DB.Model(post).Where("id = ?", entityId.ID).First(post)
	if result.Error != nil {
		logrus.Error(result.Error.Error())
		utils.InternalServerError(c, "服务错误")
		return
	}
	if result.RowsAffected == 0 {
		utils.BadRequest(c, "文章不存在")
		return
	}
	if post.UserID != middleware.CurrentUserId(c) {
		utils.Forbidden(c, "没有权限删除此文章")
		return
	}
	config.DB.Delete(post)
	utils.Success(c)
}

func (controller *PostController) Comment(c *gin.Context) {
	entityId := utils.EntityIdVO{}
	if err := c.ShouldBindUri(&entityId); err != nil {
		utils.BadRequest(c, "参数错误")
		return
	}
	vo := models.CommentCreateVO{}
	if err := c.ShouldBind(&vo); err != nil {
		utils.BadRequest(c, "参数错误")
		return
	}
	vo.PostID = entityId.ID
	vo.UserID = middleware.CurrentUserId(c)

	comment := &models.Comment{Content: vo.Content, UserID: vo.UserID, PostID: vo.PostID}
	config.DB.Create(comment)

	utils.Success(c)
}

func (controller *PostController) CommentList(c *gin.Context) {
	entityId := utils.EntityIdVO{}
	if err := c.ShouldBindUri(&entityId); err != nil {
		utils.BadRequest(c, "参数错误")
		return
	}
	var comments []models.Comment
	result := config.DB.Where("post_id = ?", entityId.ID).Find(&comments)
	if result.Error != nil {
		logrus.Error(result.Error.Error())
		utils.InternalServerError(c, "服务错误")
		return
	}
	utils.Data(c, comments)
}
