/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:20:00
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-12-30 16:11:15
 * @Email: 1321510155@qq.com
 */

package controllers

type ApiRes struct {
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type LoginParams struct {
	Username string `form:"nickName" json:"nickName" binding:"required"`
	Password string `form:"password" json:"password" bingding:"required"`
}
