
/*
 * @Author: QTTQ
 * @Date: 2018-10-25 13:25:44
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-12-30 16:30:04
 * @Email: 1321510155@qq.com
 */

 package models

 import (
	 "paopao/db"
 )
 
 type ReportForm struct {
	 Phone string `gorm:"type:varchar(100)" form:"phone" json:"phone"`
	 Symptom string `gorm:"type:varchar(100)" form:"symptom" json:"symptom"`
 }
 func SaveReportForm(phone string,symptom string) (*ReportForm, error) {
	 data := &ReportForm{Phone:phone,Symptom:symptom}
	que := db.DB.Save(&data)
	 if que.Error != nil {
		 // panic(que.Error)
		 return nil, err
	 }
	 return data, nil
 }