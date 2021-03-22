package main

import (
	"github.com/bluepongo/lifeboat_go_backend/drivers"
	"github.com/bluepongo/lifeboat_go_backend/server"
	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func main() {
	// 服务停止时清理数据库链接
	defer drivers.MysqlDb.Close()

	// 启动服务
	server.Run(HttpServer)
}
