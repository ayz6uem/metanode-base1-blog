package ui

import (
	"base1-blog/app"
	"base1-blog/infrastructure/web"

	"github.com/gin-gonic/gin"
)

func PostController(router *gin.Engine) {
	router.GET("/post", page)
	router.POST("/post", app.Logined(), create)
	router.PUT("/post/:id", app.Logined(), update)
	router.DELETE("/post/:id", app.Logined(), del)
	router.POST("/post/:id/comment", app.Logined(), comment)
	router.GET("/post/:id/comment", commentList)
}

func create(c *gin.Context) {
	vo := app.PostCreateVO{}
	err := c.ShouldBind(&vo)
	if err != nil {
		panic(err)
		return
	}
	vo.UserID = c.GetUint("OperatorId")
	app.PostCreate(vo)
	c.JSON(200, gin.H{"code": 0, "msg": "ok"})
}

func page(c *gin.Context) {
	query := web.PageQuery{}
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
		return
	}
	content, total := app.PostPage(query.Page, query.PageSize)
	c.JSON(200, gin.H{"code": 0, "msg": "ok", "data": content, "total": total})
}

func update(c *gin.Context) {
	entityId := web.EntityIdVO{}
	vo := app.PostUpdateVO{}
	err := c.ShouldBindUri(&entityId)
	if err != nil {
		panic(err)
		return
	}
	err = c.ShouldBind(&vo)
	if err != nil {
		panic(err)
		return
	}
	vo.ID = entityId.ID
	vo.UserID = c.GetUint("OperatorId")
	app.PostUpdate(vo)
	c.JSON(200, gin.H{"code": 0, "msg": "ok"})
}

func del(c *gin.Context) {
	entityId := web.EntityIdVO{}
	err := c.ShouldBindUri(&entityId)
	if err != nil {
		panic(err)
		return
	}
	app.PostDelete(entityId.ID, c.GetUint("OperatorId"))
	c.JSON(200, gin.H{"code": 0, "msg": "ok"})
}

func comment(c *gin.Context) {
	entityId := web.EntityIdVO{}
	err := c.ShouldBindUri(&entityId)
	if err != nil {
		panic(err)
		return
	}
	vo := app.CommentCreateVO{}
	err = c.ShouldBind(&vo)
	if err != nil {
		panic(err)
		return
	}
	vo.PostID = entityId.ID
	vo.UserID = c.GetUint("OperatorId")
	app.CommentCreate(vo)
	c.JSON(200, gin.H{"code": 0, "msg": "ok"})
}

func commentList(c *gin.Context) {
	entityId := web.EntityIdVO{}
	err := c.ShouldBindUri(&entityId)
	if err != nil {
		panic(err)
		return
	}
	content := app.CommentList(entityId.ID)
	c.JSON(200, gin.H{"code": 0, "msg": "ok", "content": content})
}
