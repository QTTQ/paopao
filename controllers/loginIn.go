/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:19:50
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-11-03 16:52:50
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

func GetUser(c *gin.Context) {
	uinfoIr, _ := c.Get("udata")
	uidIr, _ := c.Get("uid")
	token, err := utils.Encrypt(fmt.Sprintf("%d:%d", uidIr, time.Now().Unix()+config.TOKEN_EXPIRE_TIME), []byte(config.EncryptKey))
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "生成token失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "登录成功",
		Data: gin.H{
			"token": token,
			"user":  uinfoIr.(*models.User),
		},
	})
	return
}
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
