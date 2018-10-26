// Package main provides ...
package main

import (
	 "paopaoServer/db"
	"paopaoServer/config"
	"paopaoServer/models"
	"paopaoServer/routers"
)
func main() {
	db.ConnectAndInit(
		config.Conf,
		new(models.User),
		new(models.Article),
	)
	defer db.DB.Close()
	router := routers.InitRouters()
	router.Run(":8080")
}
