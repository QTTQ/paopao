// Package main provides ...
package main

import (
	 "paopao/db"
	"paopao/config"
	"paopao/models"
	"paopao/routers"
	"net/http"
	"time"
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
	// router.Run(":8080") //用gin 路由启动 ----
	s := &http.Server{
        Addr:           ":8080",
        Handler:        router,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe() // listen and serve on 0.0.0.0:8080  //用http  启动 ----
}
