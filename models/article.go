/*
 * @Author: QTTQ
 * @Date: 2018-10-25 13:25:44
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-11-02 12:59:58
 * @Email: 1321510155@qq.com
 */

package models

import (
	"paopao/db"
	"time"
)

var err error

type Article struct {
	Id      int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`              //id
	Uid     int    `gorm:"type:int(10)" form:"uid" json:"uid"`              //uid
	Name    string `gorm:"type:char(100)" form:"name" json:"name"`          //名字
	Actor   string `gorm:"type:varchar(800)" form:"actor" json:"actor"`     //名字
	Sex     int    `gorm:"type:int(10)" form:"sex" json:"sex"`              //性别
	Title   string `gorm:"type:varchar(100)" form:"title" json:"title"`     //文章主题
	Context string `gorm:"type:varchar(200)" form:"context" json:"context"` //文章内容
	DownloadAddress string `gorm:"type:varchar(200)" form:"downloadAddress" json:"downloadAddress"` //下载地址
	Thunmbs int    `gorm:"type:int(20)" form:"thunmbs" json:"thunmbs"`      //点赞
	CtTime  string `gorm:"type: datetime" form:"ctTime" json:"ctTime"`      //创建时间
}

//所有
func AllArticle(page int) ([]Article, error) {
	articles := []Article{} 
	// que := db.DB.Offset((page - 1) * 10+1).Limit(page * 10).Order("ct_time desc").Find(&articles)
	que := db.DB.Order("ct_time desc").Offset((page-1)*10).Limit(10).Find(&articles)
	if que.Error != nil {
		// panic(que.Error)
		return nil, que.Error
	}
	return articles, err
}

//我的文章
func MyArticle(uid, page int) ([]Article, error) {
	articles := []Article{} //获取切片   这样可以获取多个   要不 只能获取最后一个
	println(uid, page, "--我的文--------章")
	que := db.DB.Where("uid = ?", uid).Order("ct_time desc").Offset((page-1)*10).Limit(10).Find(&articles)
	if que.Error != nil {
		return nil, que.Error
	}
	println(articles, len(articles), "--我的文章")
	return articles, err
}

//写文章
func CreatArticle(uid int, title, context,paths string) (*Article, error) {
	user, err := GetUser(uid)
	if user.Actor == "" {
		user.Actor="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1540978198179&di=06da2d96bee5e89e99d376a1dac465dd&imgtype=0&src=http%3A%2F%2Fs16.sinaimg.cn%2Fmiddle%2F6b260c28g9093c271d9bf%26690"
	}
	article := Article{Uid: uid, Name: user.UserName, Actor: user.Actor, Sex: user.Sex, Title: title, Context: context,DownloadAddress:paths, CtTime: time.Now().Format("2006-01-02 15:04:05")} //获取切片   这样可以获取多个   要不 只能获取最后一个
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
func UpdateArticle(id int, text string) (*Article, error) {
	article := Article{Id: id}
	que := db.DB.Model(&article).Update("context", text)
	if que.Error != nil {
		return nil, que.Error
	}
	return &article, err
}

//点赞文章
func ThunmbToArticle(articleId int) (*Article, error) {
	article := Article{Id: articleId}
	que := db.DB.Model(&article).First(&article).Update("thunmbs", article.Thunmbs+1)
	if que.Error != nil {
		return nil, que.Error
	}
	return &article, nil
}



//点赞文章  按点赞数量最多到少 排序
func GetMostThunmbArticle(page int) ([]Article, error) {
	articles := []Article{} //获取切片   这样可以获取多个   要不 只能获取最后一个
	que := db.DB.Order("thunmbs desc").Offset((page - 1) * 10).Limit(page * 10).Find(&articles)
	if que.Error != nil {
		// panic(que.Error)
		return nil, que.Error
	}
	return articles, err
}