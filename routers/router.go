// Package routers provides ...
package routers

import (
	"paopaoServer/controllers"
	"paopaoServer/middleware"
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
	// jwtrouter := router.Group("/jwt")
	// jwtrouter.Use(middleware.UserAuth())
	// jwtrouter := router.Group("/jwt",middleware.UserAuth()) //token
	jwtrouter := router.Group("/jwt") //token
	{
		jwtrouter.POST("/AllArticle", controllers.AllArticle)
		jwtrouter.POST("/MyArticle", controllers.MyArticle)
		jwtrouter.POST("/CreatArticle", controllers.CreatArticle)
		jwtrouter.POST("/DeleteArticle", controllers.DeleteArticle)
		jwtrouter.POST("/UpdateArticle", controllers.UpdateArticle)
	}
	return router
}
