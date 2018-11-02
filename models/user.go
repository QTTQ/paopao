/*
 * @Author: QTTQ
 * @Date: 2018-10-25 13:25:44
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-11-02 13:10:29
 * @Email: 1321510155@qq.com
 */

package models

import (
	"paopao/db"
	"time"
)

type User struct {
	Uid      int    `gorm:"primary_key" form:"uid" json:"uid" `
	Unionid  int    `gorm:"type:int(20) default null" form:"unionid" json:"unionid"`
	PhoneNum int    `gorm:"type:int(20) default null" form:"phoneNum" json:"phoneNum"`
	UserName string `gorm:"type:varchar(100)" form:"name" json:"name"`
	PassWord string  `gorm:"type:varchar(100)" form:"password" json:"password"`
	Actor    string `gorm:"type:varchar(200) default null" form:"actor" json:"actor"`
	Sex      int    `gorm:"type:int(2) default 1" form:"sex" json:"sex"`
	RegTime  string `gorm:"type: datetime" form:"regTime" json:"regTime"`
}

func UserRegister(name, pass string) (*User, error) {
	t := time.Now()
	// user := User{UserName: name, PassWord: pass, RegTime: t.Format("2006-01-02 15:04:05")}
	println(pass,"-------------------------------")
	user := User{UserName: name, PassWord: pass, Sex: 1, RegTime: t.Format("2006-01-02 15:04:05")}
	db.DB.Create(&user)
	user.PassWord=""
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

func GetUser(uid int) (*User, error) {
	user := User{}
	que := db.DB.Where("uid = ?", uid).Find(&user)
	if que.Error != nil {
		// panic(que.Error)
		return nil, err
	}
	return &user, nil
}

func GetName(name string) bool {
	user := User{}
	if err := db.DB.Where("user_name = ?", name).Find(&user).Error; err != nil {
		// panic(que.Error)
		return false
	}
	if len(user.UserName) != 0 {
		return true
	}
	return false
}
