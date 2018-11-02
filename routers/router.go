// Package routers provides ...
package routers

import (
	"paopao/controllers"
	"paopao/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// router.LoadHTMLGlob("templates/*")
	// router.Static("/assets", "./assets")
	
	router.Use(middleware.Cors())
	router.POST("/Register", controllers.Register)// 注册  
	router.POST("/LoginIn", controllers.LoginIn)  //登录 
	router.POST("/AllArticle", controllers.AllArticle) //获取所有文章
	router.POST("/GetMostThunmbArticle", controllers.GetMostThunmbArticle)//获取点赞最多文章
	// jwtrouter := router.Group("/jwt")
	// jwtrouter.Use(middleware.UserAuth())
	jwtrouter := router.Group("/jwt",middleware.UserAuth()) //token
	// jwtrouter := router.Group("/jwt") //token
	{
		jwtrouter.POST("/MyArticle", controllers.MyArticle)

		//文章
		jwtrouter.POST("/CreatArticle", controllers.CreatArticle)
		jwtrouter.POST("/DeleteArticle", controllers.DeleteArticle)
		jwtrouter.POST("/UpdateArticle", controllers.UpdateArticle)
		jwtrouter.POST("/ThunmbToArticle", controllers.ThunmbToArticle)
		
		//评论
		jwtrouter.POST("/CurrentArticledMessages", controllers.CurrentArticledMessages)
		jwtrouter.POST("/CurrentArticledWirteMessage", controllers.CurrentArticledWirteMessage)
		jwtrouter.POST("/CurrentArticledMessageOtherMessage", controllers.CurrentArticledMessageOtherMessage)
		jwtrouter.POST("/ThunmbToCurrentArticledMessage", controllers.ThunmbToCurrentArticledMessage)


		//upload
		jwtrouter.POST("/Upload", controllers.Upload)

	}
	return router
}
