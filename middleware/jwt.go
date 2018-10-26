/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:20:13
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-10-23 15:53:15
 * @Email: 1321510155@qq.com
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paopaoServer/config"
	"paopaoServer/controllers"
	"paopaoServer/models"
	"paopaoServer/util"
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
						uinfo,err := models.GetUser(uid)
						if err!=nil{
							panic(err)
						}
						c.Set(USER_ID_KEY, uid)
						c.Set(USER_DATA_KEY, uinfo)
						return
					}

				}
			}
		}
		c.Abort()
		c.JSON(http.StatusUnauthorized,controllers.ApiRes{
			Code:1,
			Msg:"用户未登录",
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
