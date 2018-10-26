/*
 * @Author: QTTQ
 * @Date: 2018-10-25 13:25:44
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-10-25 16:51:58
 * @Email: 1321510155@qq.com
 */

package models

import (
	"paopaoServer/db"
	"time"
)

var err error

const (
	PAGE_SIZE = 2
)

type Article struct {
	Id      int    `gorm:"AUTO_INCREMENT"`    //id
	Uid     int    `gorm:"type:int(10)"`      //uid
	Name   string `gorm:"type:char(100)"` //名字
	Actor   string `gorm:"type:varchar(200)"` //名字
	Sex   int `gorm:"type:int(10)"` //性别
	Title   string `gorm:"type:varchar(100)"` //文章主题
	Context string `gorm:"type:varchar(100)"` //文章内容
	CtTime  string `gorm:"type: datetime"`    //创建时间
}

//所有
func AllArticle(page int) ([]Article, error) {
	articles := []Article{} //获取切片   这样可以获取多个   要不 只能获取最后一个
	que := db.DB.Order("ct_time desc").Offset((page - 1) * 10).Limit(page * 10).Find(&articles)
	if que.Error != nil {
		// panic(que.Error)
		return nil, que.Error
	}
	return articles, err
}

//我的文章
func MyArticle(uid, page int) ([]Article, error) {
	articles := []Article{} //获取切片   这样可以获取多个   要不 只能获取最后一个
	println(uid, page , "--我的文--------章")
	que := db.DB.Where("uid = ?", uid).Order("ct_time desc").Offset((page - 1) * 10).Limit(page * 10).Find(&articles)
	if que.Error != nil {
		return nil, que.Error
	}
	println(articles, len(articles), "--我的文章")
	return articles, err
}

//写文章
func CreatArticle(uid int, title, context string) (*Article, error) {
user,err:=GetUser(uid)
	article := Article{Uid: uid,Name:user.UserName,Actor:user.Actor,Sex:user.Sex, Title: title, Context: context, CtTime: time.Now().Format("2006-01-02 15:04:05")} //获取切片   这样可以获取多个   要不 只能获取最后一个
	que := db.DB.Create(&article)
	if que.Error != nil {
		// panic(que.Error)
		return nil, que.Error
	}
	return &article, err
}
//删除写文章
func DeleteArticle(id int) (*Article, error) {
	article := Article{Id: id} 
	que := db.DB.Delete(&article)
	if que.Error != nil {
		return nil, que.Error
	}
	return &article, err
}

//更新文章
func UpdateArticle(id int,text string) (*Article, error) {
	article := Article{Id:id} 
	que := db.DB.Model(&article).Update("context",text)
	if que.Error != nil {
		return nil, que.Error
	}
	return &article, err
}
