// Package main provides ...
package main

import (
	 "paopao/db"
	"paopao/config"
	"paopao/models"
	"paopao/routers"
)
func main() {
	db.ConnectAndInit(
		config.Conf,
		new(models.User),
		new(models.Article),
		new(models.Message),
	)
	defer db.DB.Close()
	router := routers.InitRouters()
	router.Run(":8080")
}
