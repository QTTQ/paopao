/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:20:00
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-10-25 14:02:11
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
	regParams := LoginParams{}
	err := c.Bind(&regParams)
	if regParams.Password == "" || regParams.Username == "" {
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
	if len(regParams.Username) < 0 || len(regParams.Password) < 0 {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "账号或密码不能为空",
			})
		return
	}
	hadUser := models.GetName(regParams.Username) //判断是否已经注册
	if hadUser {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "用户已经存在",
			})
		return
	}
	user, err := models.UserRegister(regParams.Username, regParams.Password)
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
