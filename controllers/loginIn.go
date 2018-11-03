/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:19:50
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-11-03 14:37:52
 * @Email: 1321510155@qq.com
 */

package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"paopao/config"
	"paopao/models"
	"paopao/util"
	// "strconv"
	"time"
)

func LoginIn(c *gin.Context) {
	loginParams := models.User{}
	err := c.Bind(&loginParams)
	if len(loginParams.UserName) <= 0 || len(loginParams.PassWord) <= 0 {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "账号或密码不能为空",
		})
	}
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "登录数据格式不正确！",
		})
		return
	}
	user, err := models.UserLogin(loginParams.UserName, loginParams.PassWord)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "登录失败",
		})
		return
	}
	token, err := utils.Encrypt(fmt.Sprintf("%d:%d", user.Uid, time.Now().Unix()+config.TOKEN_EXPIRE_TIME), []byte(config.EncryptKey))
	// type UserData struct {
	// 	Uid      int    `json:"uid"`
	// 	PhoneNum int    `json:"phoneNum"`
	// 	UserName string `json:"name"`
	// 	Actor    string `json:"actor"`
	// 	Sex      int    `json:"sex"`
	// }

	// // var	userData map[string]string
	// var userData map[string]string = map[string]string{}
	// // var	userData map[string]string =make(map[string]string,10)
	// userData["name"] = user.UserName
	// userData["phoneNum"] = strconv.Itoa(user.PhoneNum)
	// userData["actor"] = user.Actor
	// userData["uid"] = strconv.Itoa(user.Uid)
	// userData["sex"] = strconv.Itoa(user.Sex)

	// userData := UserData{}
	// userData.Uid = user.Uid
	// userData.UserName = user.UserName
	// userData.PhoneNum = user.PhoneNum
	// userData.Actor = user.Actor
	// userData.Sex = user.Sex
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "登录成功",
		Data: gin.H{
			"token": token,
			// "user":  &userData,
			"user": user,
		},
	})
	return
}
