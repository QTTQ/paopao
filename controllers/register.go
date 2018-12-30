/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:20:00
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-12-29 11:02:34
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
	"time"
	// "io/ioutil"
)

func Register(c *gin.Context) {
	regParams := models.User{}
	err := c.Bind(&regParams)
	fmt.Println(regParams, "--------------------------")
	if regParams.PassWord == "" || regParams.UserName == "" {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "姓名或密码不能为空",
			})
		return
	}
	if err != nil {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "获取数据错误",
			})
		return
	}
	if len(regParams.UserName) < 0 || len(regParams.PassWord) < 0 {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "账号或密码不能为空",
			})
		return
	}
	hadUser := models.GetName(regParams.UserName) //判断是否已经注册
	if hadUser {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "用户已经存在",
			})
		return
	}
	user, err := models.UserRegister(regParams.UserName, regParams.PassWord, regParams.Actor, regParams.Sex)
	if err != nil {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "登录数据格式不正确！",
			})
		return
	}
	token, err := utils.Encrypt(fmt.Sprintf("%d:%d", user.Uid, time.Now().Unix()+config.TOKEN_EXPIRE_TIME), []byte(config.EncryptKey))
	c.JSON(http.StatusOK,
		ApiRes{
			Code: 0,
			Msg:  "成功注册",
			Data: gin.H{
				"token": token,
			},
		})
	return
}
