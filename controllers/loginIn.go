/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:19:50
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-10-25 14:34:45
 * @Email: 1321510155@qq.com
 */

package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"paopaoServer/config"
	"paopaoServer/models"
	"paopaoServer/util"
	"time"
)

func LoginIn(c *gin.Context) {
	loginParams := LoginParams{}
	err := c.Bind(&loginParams)
	if len(loginParams.Username) <= 0 || len(loginParams.Password) <= 0 {
		c.JSON(200, ApiRes{
			Code: 1,
			Msg:  "账号或密码不能为空",
		})
	}
	if err != nil {
		c.JSON(200, ApiRes{
			Code: 1,
			Msg:  "登录数据格式不正确！",
		})
		return
	}
	user, err := models.UserLogin(loginParams.Username, loginParams.Password)
	if err != nil {
		c.JSON(200, ApiRes{
			Code: 1,
			Msg:  "登录失败",
		})
		return
	}
	token, err := utils.Encrypt(fmt.Sprintf("%d:%d", user.Uid, time.Now().Unix()+config.TOKEN_EXPIRE_TIME), []byte(config.EncryptKey))
	type UserData struct {
		Uid      int   `gorm:"primary_key"`
		UserName string `json:"username"`
		Actor string `gorm:"default null"`
		Sex int `gorm:"default 1"`
		RegTime  string
	}
	data:=UserData{}
	data.UserName=user.UserName
	data.Actor=user.Actor
	data.Uid=user.Uid
	data.Sex=user.Sex
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "登录成功",
		Data: gin.H{
			"token": token,
			"user":&data,
		},
	})
	return
}

// func Async(c *gin.Context) {
// 	cCp := c.Copy()
// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		log.Println("Done! in path" + cCp.Request.URL.Path)
// 	}()
// 	c.JSON(http.StatusOK, gin.H{"data": "Async"})
// }
