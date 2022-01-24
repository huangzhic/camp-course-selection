package server

import (
	"camp-course-selection/api"
	"camp-course-selection/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	g := r.Group("/api/v1")

	g.POST("/auth/login", api.Login)
	auth := g.Group("")
	auth.Use(middleware.AuthRequired())
	{
		// 成员管理
		auth.GET("/member")
		auth.POST("/member/create", api.CreateMember)
		auth.GET("/member/list")
		auth.POST("/member/update")
		auth.POST("/member/delete")

		// 登录
		auth.POST("/auth/logout", api.Logout)

		auth.GET("/auth/whoami", api.Whoami)

		// 排课
		auth.POST("/course/create")
		auth.GET("/course/get")

		auth.POST("/teacher/bind_course")
		auth.POST("/teacher/unbind_course")
		auth.GET("/teacher/get_course")
		auth.POST("/course/schedule")

		// 抢课
		auth.POST("/student/book_course")
		auth.GET("/student/course")
	}
	return r
}
