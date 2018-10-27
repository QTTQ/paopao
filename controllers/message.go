/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:20:00
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-10-25 14:02:11
 * @Email: 1321510155@qq.com
 */

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paopao/models"
	"strconv"
	// "io/ioutil"
)

type MessageParams struct {
Id  string `form:"id" json:"id"`
	ArticleId  string `form:"articleId" json:"articleId"`
	Uid        string `form:"uid" json:"uid"`
	UserName   string `form:"name" json:"name"`
	Actor   string `form:"actor" json:"actor"`
	Context    string `form:"context" json:"context"`
	Thunmbs    string `form:"thunmbs" json:"thunmbs"`
	ThunmbsUid string `form:"thunmbsUid" json:"thunmbsUid"`
	ToMesId      string    `form:"toMesId" json:"toMesId"`
	Page       string `form:"page" json:"page"`
}

func CurrentArticledMessages(c *gin.Context) {
	mesParams := MessageParams{}
	err := c.Bind(&mesParams)
	artId, err := strconv.Atoi(mesParams.ArticleId)
	if artId <= 0 || err != nil {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "articleId参数错误",
			})
		return
	}
	page, err := strconv.Atoi(mesParams.Page)
	if page <= 0 || err != nil {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "page参数错误",
			})
		return
	}
	messages, err := models.CurrentArticledMessages(artId, page)
	if err != nil {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "获取所有评论请求数据错误",
			})
		return
	}
	c.JSON(http.StatusOK,
		ApiRes{
			Code: 0,
			Msg:  "成功",
			Data: gin.H{
				"messages": messages,
			},
		})
	return
}

func CurrentArticledWirteMessage(c *gin.Context) {
	mesParams := MessageParams{}
	err := c.Bind(&mesParams)
	artId, err := strconv.Atoi(mesParams.ArticleId)
	if artId <= 0 || err != nil {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "参数错误",
			})
		return
	}
	mesUid, err := strconv.Atoi(mesParams.Uid)
	if mesUid <= 0 || err != nil {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "参数错误",
			})
		return
	}
	message, err := models.CurrentArticledWirteMessage(artId, mesUid, mesParams.UserName, mesParams.Actor, mesParams.Context)
	if err != nil {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "评论请求数据错误",
			})
		return
	}
	c.JSON(http.StatusOK,
		ApiRes{
			Code: 0,
			Msg:  "成功",
			Data: gin.H{
				"message": message,
			},
		})
	return
}

func CurrentArticledMessageOtherMessage(c *gin.Context) {
	mesParams := MessageParams{}
	err := c.Bind(&mesParams)
	artId, err := strconv.Atoi(mesParams.ArticleId)
	if artId <= 0 || err != nil {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "参数错误",
			})
		return
	}
	mesUid, err := strconv.Atoi(mesParams.Uid)
	if mesUid <= 0 || err != nil {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "参数错误",
			})
		return
	}
	toMesId, err := strconv.Atoi(mesParams.ToMesId)
	if toMesId <= 0 || err != nil {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "参数错误",
			})
		return
	}
	message, err := models.CurrentArticledMessageOtherMessage(artId, mesUid, toMesId, mesParams.UserName, mesParams.Context)
	if err != nil {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "评论其他人请求数据错误",
			})
		return
	}
	c.JSON(http.StatusOK,
		ApiRes{
			Code: 0,
			Msg:  "成功",
			Data: gin.H{
				"message": message,
			},
		})
	return
}

func ThunmbToCurrentArticledMessage(c *gin.Context) {
	mesParams := MessageParams{}
	err := c.Bind(&mesParams)
	artId, err := strconv.Atoi(mesParams.Id)
	if artId <= 0 || err != nil {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "参数错误",
			})
		return
	}
	message, err := models.ThunmbToCurrentArticledMessage(artId)
	if err != nil {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "点赞当前请求数据错误",
			})
		return
	}
	c.JSON(http.StatusOK,
		ApiRes{
			Code: 0,
			Msg:  "成功",
			Data: gin.H{
				"message": message,
			},
		})
	return
}
