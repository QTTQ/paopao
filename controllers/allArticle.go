/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:19:50
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-11-02 20:14:28
 * @Email: 1321510155@qq.com
 */

package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"paopao/models"
	"strconv"
)

const DST = "static/img/"

func AllArticle(c *gin.Context) {
	// page:=c.PostForm("page")
	page := c.DefaultPostForm("page", "0")
	n, err := strconv.Atoi(page)
	allarticle, err := models.AllArticle(n)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "登录失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "获取文章成功",
		Data: gin.H{
			"data": allarticle,
		},
	})
	return
}

type MyArticleParams struct {
	Uid  string `form:"uid" json:"uid"`
	Page string `form:"page" json:"page"`
}

func MyArticle(c *gin.Context) {
	myParams := MyArticleParams{}
	c.Bind(&myParams)
	uidInter, _ := c.Get("uid")
	uinfoInter, _ := c.Get("udata")

	uid := uidInter.(int)//取接口里的type类型
	uininfo:=uinfoInter.(*models.User)
	fmt.Println(uid,uininfo.Uid,"---------------uininfo-----------------") //取接口里的结构体
	page, err := strconv.Atoi(myParams.Page)
	allarticle, err := models.MyArticle(uid, page)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "登录失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "获取文章成功",
		Data: gin.H{
			"data": allarticle,
		},
	})
	return
}

type CreatArticleParams struct {
	Uid     string `form:"uid" json:"uid"`         //uid
	Title   string `form:"title" json:"title"`     //文章主题
	Context string `form:"context" json:"context"` //文章内容
}

func CreatArticle(c *gin.Context) {
	createParams := CreatArticleParams{}
	c.Bind(&createParams)
	uid, err := strconv.Atoi(createParams.Uid)
	paths := ""
	form, _ := c.MultipartForm()
	files := form.File["upload"]

	for _, file := range files {
		paths += DST + file.Filename + ","
		// Upload the file to specific dst.
		c.SaveUploadedFile(file, DST+file.Filename)
	}
	article, err := models.CreatArticle(uid, createParams.Title, createParams.Context, paths)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "登录失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "创建文章成功",
		Data: gin.H{
			"data": article,
		},
	})
	return
}
func DeleteArticle(c *gin.Context) {
	deleteParams := MyArticleParams{}
	c.Bind(&deleteParams)
	page, err := strconv.Atoi(deleteParams.Page)
	res, err := models.DeleteArticle(page)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "删除文章失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "删除文章成功",
		Data: gin.H{
			"data": res,
		},
	})
	return
}

//删除
type UpdateArticleParams struct {
	Id      string `form:"id" json:"id"`
	Context string `form:"text" json:"text"`
}

func UpdateArticle(c *gin.Context) {
	updateParams := UpdateArticleParams{}
	c.Bind(&updateParams)
	id, err := strconv.Atoi(updateParams.Id)
	res, err := models.UpdateArticle(id, updateParams.Context)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "删除文章失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "删除文章成功",
		Data: gin.H{
			"data": res,
		},
	})
	return
}

func ThunmbToArticle(c *gin.Context) {
	artIdStr := c.DefaultPostForm("artId", "0")
	artId, err := strconv.Atoi(artIdStr)
	println(artId, "----------------------")
	if err != nil || artId <= 0 {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "参数错误",
		})
		return
	}
	res, err := models.ThunmbToArticle(artId)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "点赞失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "点赞成功",
		Data: gin.H{
			"data": res,
		},
	})
	return
}

func GetMostThunmbArticle(c *gin.Context) {
	page := c.DefaultPostForm("page", "0")
	n, err := strconv.Atoi(page)
	allarticle, err := models.GetMostThunmbArticle(n)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "获取点赞文章失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "获取文章成功",
		Data: gin.H{
			"data": allarticle,
		},
	})
	return
}
