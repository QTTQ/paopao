/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:19:50
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-11-28 20:15:31
 * @Email: 1321510155@qq.com
 */

package controllers

import (
	"context"
	"fmt"
	"github.com/akkuman/parseConfig"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"net/http"
	"os"
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
	uidIr, _ := c.Get("uid")
	uinfoIr, _ := c.Get("udata")

	uid := uidIr.(int) //取接口里的type类型
	uininfo := uinfoIr.(*models.User)
	fmt.Println(uid, uininfo.Uid, "---------------uininfo-----------------") //取接口里的结构体
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
	Title   string `form:"title" json:"title"`     //文章主题
	Context string `form:"context" json:"context"` //文章内容
}

func uploadQiNiu(localFile string) {
	//取出配置文件
	config := parseConfig.New("config/conf.json")
	//七牛云配置
	// downLoadDirPath := "./"
	accessKey := config.Get("qiniu_config > accessKey").(string)
	secretKey := config.Get("qiniu_config > secretKey").(string)
	bucket := config.Get("qiniu_config > bucket").(string)
	fmt.Println(accessKey, secretKey, bucket, "-----mmmmm-----")
	// videoDomain := config.Get("qiniu_config > url").(string)

	mac := qbox.NewMac(accessKey, secretKey)
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(mac)
	fmt.Println(upToken)
	cfg := storage.Config{}
	//华北机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// key := "github-x.png"
	key :=localFile
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}

func CreatArticle(c *gin.Context) {
	// ua := c.GetHeader("User-Agent")
	// ct := c.GetHeader("Content-Type")
	// fmt.Println(ua, "-ua-", "\n", ct, "--ct--")

	createParams := CreatArticleParams{}
	c.Bind(&createParams)
	uidIr, _ := c.Get("uid")
	uid := uidIr.(int)
	paths := ""
	form, _ := c.MultipartForm()
	files := form.File["file"]
	fmt.Println(form, files, "--------------------")
	for _, file := range files {
		paths += DST + file.Filename + ","
		// Upload the file to specific dst.
		c.SaveUploadedFile(file, DST+file.Filename)
		uploadQiNiu(DST + file.Filename) //上传七牛
		err := os.Remove(DST + file.Filename)
		if err != nil {
			fmt.Println(err)
		}
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

// type CreatArticleParams struct {
// 	// Uid     string `form:"uid" json:"uid"`         //uid
// 	Title   string   `form:"title" json:"title"`     //文章主题
// 	Context string   `form:"context" json:"context"` //文章内容
// 	Paths   []string `form:"paths" json:"paths"`     //地址

// }

// func CreatArticle(c *gin.Context) {
// 	fmt.Println("===============111111111111====================")

// 	createParams := CreatArticleParams{}
// 	c.Bind(&createParams)
// 	paths := ""
// 	fmt.Println(createParams.Context, "===================================")
// 	uidIr, _ := c.Get("uid")

// 	uid := uidIr.(int)
// 	c.JSON(http.StatusOK, ApiRes{
// 		Code: 1,
// 		Msg:  "获取uid失败",
// 	})
// 	for _, file := range createParams.Paths {
// 		paths += file + ","
// 	}
// 	article, err := models.CreatArticle(uid, createParams.Title, createParams.Context, paths)
// 	if err != nil {
// 		c.JSON(http.StatusOK, ApiRes{
// 			Code: 1,
// 			Msg:  "登录失败",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, ApiRes{
// 		Code: 0,
// 		Msg:  "创建文章成功",
// 		Data: gin.H{
// 			"data": article,
// 		},
// 	})
// 	return
// }

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
	form, _ := c.MultipartForm()
	artId, err := strconv.Atoi(artIdStr)
	fmt.Println(artId, "----", form, "----------------------")
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
