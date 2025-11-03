package routes

import (
	"base1-blog/controllers"
	"base1-blog/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()

	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.ErrorHandlingMiddleware())
	r.Use(gin.Recovery())

	authController := &controllers.AuthController{}
	userController := &controllers.UserController{}
	postController := &controllers.PostController{}

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}
		user := api.Group("/users")
		user.Use(middleware.AuthMiddleware())
		{
			user.GET("/info", userController.Info)
		}
		post := api.Group("/posts")
		post.Use(middleware.AuthMiddleware())
		{
			post.POST("", postController.Create)
			post.PUT("/:id", postController.Update)
			post.DELETE("/:id", postController.Del)
			post.POST("/:id/comments", postController.Comment)
		}
		publicPost := api.Group("/posts")
		{
			publicPost.GET("", postController.Page)
			publicPost.GET("/:id/comments", postController.CommentList)
		}
	}
	return r
}
