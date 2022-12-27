package main

import (
	"oa-backend/models"
	"oa-backend/routes"
)

func main() {
	//数据库初始化
	models.Init()
	//路由初始化
	routes.Init()
}
