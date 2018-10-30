/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:19:50
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-10-25 16:47:53
 * @Email: 1321510155@qq.com
 */

package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	//  "paopao/models"
	//  "strconv"
	"io"
	"log"
)

func Upload(c *gin.Context) {

	name := c.PostForm("name")
	fmt.Println(name)
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	filename := header.Filename

	fmt.Println(file, err, filename)

	out, err := os.Create("static/img/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusCreated, "upload successful")

	filepath := "http://127.0.0.1:8000/file/" + filename
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "获取文章成功",
		Data: gin.H{
			"filepath": filepath,
		},
	})
	return
}
