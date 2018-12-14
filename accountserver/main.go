package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"xiaolehuigo/accountserver/accountrpc"
	"xiaolehuigo/accountserver/config"
	"xiaolehuigo/accountserver/middlewares"
	"xiaolehuigo/accountserver/router"
)

func main() {
	channel := make(chan int)
	go accountrpc.Listen(channel)
	<-channel
	config.Init()
	ginRouter := gin.New()
	ginRouter.Use(middlewares.RecoverMiddleWares())
	ginRouter.Use(middlewares.LoggerMiddlerware())
	router.ConfigRouter(ginRouter)
	ginRouter.Run(":8080")
}
