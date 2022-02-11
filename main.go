package main

import (
	"camp-course-selection/conf"
	"camp-course-selection/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":4000")
}
