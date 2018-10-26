/*
 * @Author: QTTQ
 * @Date: 2018-10-25 13:25:44
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-10-25 16:50:23
 * @Email: 1321510155@qq.com
 */

 package models

 import (
	 "paopaoServer/db"
	 "time"
 )
 
 type Message struct {
	 Uid      uint   `gorm:"primary_key"  json:"uid"`
	 UserName string `json:"username"`
	 PassWord string `json:"password"`
	 RegTime  string
 }
 
 func MessageUserRegister(name, pass string) (*User, error) {
	 t := time.Now()
	 user := User{UserName: name, PassWord: pass, RegTime: t.Format("2006-01-02 15:04:05")}
	 db.DB.Create(&user)
	 return &user, nil
 }
 
 