
/*
 * @Author: QTTQ
 * @Date: 2018-10-25 13:25:44
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-12-29 16:06:23
 * @Email: 1321510155@qq.com
 */

 package models

 import (
	 "paopao/db"
 )
 
 type ShouzhenSelectData struct {
	 Type string `gorm:"type:varchar(100)" form:"type" json:"type"`
	 Category string `gorm:"type:varchar(100)" form:"category" json:"category"`
 }
 func GetShouzhenSelectData() ([]ShouzhenSelectData, error) {
	 data := []ShouzhenSelectData{}
	que := db.DB.Find(&data)

	 if que.Error != nil {
		 // panic(que.Error)
		 return nil, err
	 }
	 return data, nil
 }