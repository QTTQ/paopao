/*
 * @Author: QTTQ
 * @Date: 2018-10-25 13:25:44
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-10-25 16:50:23
 * @Email: 1321510155@qq.com
 */

package models

import (
	"paopao/db"
	"time"
)

type Message struct {
	Id  int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	ArticleId  int    `gorm:"type:int(20)" form:"articleId" json:"articleId"`
	Uid        int    `gorm:"type:int(20)" form:"uid" json:"uid"`
	UserName   string `gorm:"type:char(100)" form:"name" json:"name"`
	Actor   string `gorm:"type:char(100)" form:"actor" json:"actor"`
	Context    string `gorm:"type:varchar(200)" form:"context" json:"context"`
	Thunmbs    int    `gorm:"type:int(20)" form:"thunmbs" json:"thunmbs"`
	ThunmbsUid int    `gorm:"type:int(20)" form:"thunmbsUid" json:"thunmbsUid"`
	ToMesId      int    `gorm:"type:int(20)" form:"toMesId" json:"toMesId"`
	SendTime   string `gorm:"type:datetime" form:"sendTime" json:"sendTime"`
}

func CurrentArticledMessages(articleId, page int) ([]Message, error) {
	messages := []Message{}
	que := db.DB.Where("article_id = ?", articleId).Offset((page - 1) * 10).Limit(page * 10).Find(&messages)
	if que.Error != nil {
		return nil, que.Error
	}
	return messages, nil
}

func CurrentArticledWirteMessage(articleId, uid int, name, actor,cxt string) (*Message, error) {
	t := time.Now()
	message := Message{ArticleId: articleId, Uid: uid, UserName: name,Actor:actor, Context: cxt, SendTime: t.Format("2006-01-02 15:04:05")}
	que := db.DB.Create(&message)
	if que.Error != nil {
		return nil, que.Error
	}
	return &message, nil
}
func CurrentArticledMessageOtherMessage(articleId, uid, toMesId int, name, cxt string) (*Message, error) {
	t := time.Now()
	message := Message{ArticleId: articleId, Uid: uid, UserName: name, Context: cxt, ToMesId: toMesId, SendTime: t.Format("2006-01-02 15:04:05")}
	que := db.DB.Create(&message)
	if que.Error != nil {
		return nil, que.Error
	}
	return &message, nil
}

func ThunmbToCurrentArticledMessage(artId int) (*Message, error) {
	message := Message{Id: artId}
	que := db.DB.Model(&message).First(&message).Update("thunmbs", message.Thunmbs+1)
	if que.Error != nil {
		return nil, que.Error
	}
	return &message, nil
}
