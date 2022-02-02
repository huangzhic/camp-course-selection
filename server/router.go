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

	g.POST("/auth/login", api.Login) // 登录后设置了session   user_id   与CurrentUser 中的user_id相呼应
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
		auth.POST("/course/create", api.CreateCourse)
		auth.GET("/course/get", api.GetCourse)

		auth.POST("/teacher/bind_course", api.BindCourse)
		auth.POST("/teacher/unbind_course", api.UnBindCourse)
		auth.GET("/teacher/get_course", api.GetTeacherCourse)
		auth.POST("/course/schedule", api.ScheduleCourse)

		// 抢课
		auth.POST("/student/book_course")
		auth.GET("/student/course")
	}
	return r
}
