/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:20:13
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-12-03 15:01:23
 * @Email: 1321510155@qq.com
 */

 package middleware

 import (
	 "fmt"
	 "github.com/gin-gonic/gin"
 )

 func AllUrlPath(c *gin.Context){
	path:=c.Request.RequestURI// 获取完整地址栏
	path1:=c.Request.URL// 获取完整地址栏
	fmt.Println(path,"---path---获取完整地址栏---")
	fmt.Println(path1,"---path1---获取完整地址栏---")

 }