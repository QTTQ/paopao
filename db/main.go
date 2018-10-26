package db

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"paopaoServer/config"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB        *gorm.DB
	ErrDBNull = errors.New("db is null,please connect first")
)

func ConnectAndInit(conf config.ConfStruct, models ...interface{}) {
	// dbUser := conf.DB.User
	// dbPass := conf.DB.Password
	// dbHost := conf.DB.Host
	// dbName := conf.DB.Name
	// dbPort := fmt.Sprintf("%d", conf.DB.Port)
	// dbPrefix := conf.DB.Prefix
	dbUser := "root"
	dbPass := "123456"
	dbHost := "127.0.0.1"
	dbName := "paopao"
	dbPort := fmt.Sprintf("%d", 3306)
	// dbPrefix := "prefix"
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		// if defaultTableName == "meta" {
		// 	defaultTableName = "metas"
		// }
		// return dbPrefix + defaultTableName
		return defaultTableName
	}
	Connect(dbUser, dbPass, dbHost, dbPort, dbName)
	Init(models...)
}
func Init(models ...interface{}) error {
	if DB == nil {
		return ErrDBNull
	}
	DB.AutoMigrate(models...)
	return nil
}
func Connect(user, password, host string, dbPort, dbName string) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbPort, dbName))
	if err != nil {
		panic("connect database failed" + err.Error())
	}
	db.SingularTable(true) //取消建表为复数形式
	DB = db
}
