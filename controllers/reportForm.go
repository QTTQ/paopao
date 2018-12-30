
/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:19:50
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-12-30 16:32:51
 * @Email: 1321510155@qq.com
 */

 package controllers

 import (
	 "fmt"
	 "github.com/gin-gonic/gin"
	 "net/http"
	 "paopao/models"
 )
 
 func GetReportForm(c *gin.Context) {
	regParams := models.ReportForm{}
	err := c.Bind(&regParams)
	 data, err := models.SaveReportForm(regParams.Phone,regParams.Symptom)
	 fmt.Println(data, "--------------------------")
	 if err != nil {
		 c.JSON(http.StatusOK,
			 ApiRes{
				 Code: 1,
				 Msg:  "获取数据错误",
			 })
		 return
	 }
	 c.JSON(http.StatusOK, gin.H{"code": "0", "msg": "成功", "results": data})
	 return
 }
 