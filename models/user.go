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

type User struct {
	Uid      int   `gorm:"primary_key"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Actor string `gorm:"default null"`
	Sex int `gorm:"default 1"`
	RegTime  string
}

func UserRegister(name, pass string) (*User, error) {
	t := time.Now()
	// user := User{UserName: name, PassWord: pass, RegTime: t.Format("2006-01-02 15:04:05")}
	user := User{UserName: name, PassWord: pass,Actor:"",Sex:1, RegTime: t.Format("2006-01-02 15:04:05")}
	db.DB.Create(&user)
	return &user, nil
}

func UserLogin(name, pass string) (*User, error) {
	user := User{}
	var err error
	que := db.DB.Where("user_name = ? AND pass_word = ?", name, pass).First(&user)
	if que.Error != nil {
		return nil, err
	}
	if len(user.UserName) != 0 {
		return &user, nil
	}
	return nil, err
}

func GetUser(uid int) (*User ,error){
	user := User{}
	que := db.DB.Where("uid = ?", uid).Find(&user)
	if que.Error != nil {
		// panic(que.Error)
		return nil, err
	}
	return &user,nil
}

func GetName(name string) bool {
	user := User{}
	if err:=db.DB.Where("user_name = ?", name).Find(&user).Error; err != nil {
		// panic(que.Error)
	return false
	}
	if len(user.UserName) != 0 {
		return true
	}
	return false
}
