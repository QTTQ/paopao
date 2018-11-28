/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:20:13
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-11-17 13:57:40
 * @Email: 1321510155@qq.com
 */

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"paopao/config"
	"paopao/controllers"
	"paopao/models"
	"paopao/util"
	"strconv"
	"strings"
	"time"
)

const (
	AUTH_HEADER_NAME = "X-Auth-Token"
	USER_ID_KEY      = "uid"
	USER_DATA_KEY    = "udata"
	DB_ERR_CODE      = 258
)

func respond(c *gin.Context, code uint, msg string, data interface{}) {
	c.JSON(http.StatusOK, controllers.ApiRes{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(AUTH_HEADER_NAME)
		if len(token) > 0 {
			uData, err := utils.Decrypt(token, []byte(config.EncryptKey))
			if err == nil {
				parts := strings.Split(uData, ":")
				if len(parts) == 2 {
					expireTime, err := strconv.ParseInt(parts[1], 10, 64)
					uid, _ := strconv.Atoi(parts[0])
					if err == nil && expireTime > time.Now().Unix() && uid > 0 {
						uinfo, err := models.GetUser(uid)
						if err != nil {
							panic(err)
						}
						fmt.Println(uid, uinfo, "----------------成功-------")
						c.Set(USER_ID_KEY, uid)
						c.Set(USER_DATA_KEY, uinfo)
						return
					}
				}
			}
		}
		fmt.Println("---------------jwt失败-------")
		c.Abort()
		c.JSON(http.StatusUnauthorized, controllers.ApiRes{
			Code: 1,
			Msg:  "用户未登录",
		})
		return
	}
}

// func Async(c *gin.Context) {
// 	cCp := c.Copy()
// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		log.Println("Done! in path" + cCp.Request.URL.Path)
// 	}()
// 	c.JSON(http.StatusOK, gin.H{"data": "Async"})
// }
