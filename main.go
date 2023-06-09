package main

import (
	"account/handler"
	"account/log"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//web server 监听在9999端口
	ip := flag.String("ip", "127.0.0.1", "请输入ip")
	port := flag.Int64("port", 9999, "请输入端port")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *ip, *port)
	log.InitLogger()
	r := gin.Default()
	accountGroup := r.Group("/v1/account")
	{
		accountGroup.GET("/list", handler.AccountListHandler)
		accountGroup.GET("/mobile", handler.GetAccountByMobileHandler)
		accountGroup.GET("id", handler.GetAccountByIDHandler)
		accountGroup.POST("/add", handler.AddAccountHandler)
		accountGroup.POST("/update", handler.UpdateAccountHandler)
	}
	r.Run(addr)
}
