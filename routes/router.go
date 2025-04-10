package routes

import (
	"go-blog-build/controller"
	"go-blog-build/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 用户注册路由
	router.POST("/register", controller.RegisterUser)
	// 用户登录路由
	router.POST("/login", controller.LoginUser)

	// 受保护的路由组，需要认证
	protected := router.Group("/protected", middleware.AuthMiddleware())
	{
		// 文章创建路由
		protected.POST("/create-article", controller.CreateArticle)
		// 获取所有文章路由
		protected.GET("/articles", controller.GetArticles)
		// 获取单个文章路由
		protected.GET("/articles/:id", controller.GetArticle)
		// 文章更新路由
		protected.PUT("/articles/:id", controller.UpdateArticle)
		// 文章删除路由
		protected.DELETE("/articles/:id", controller.DeleteArticle)
	}

	return router
}
