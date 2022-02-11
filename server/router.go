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
<<<<<<< HEAD
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
		auth.POST("/student/book_course",api.Course)
		auth.GET("/student/course", api.BookCourse)
=======

	// 成员管理
	g.GET("/member", api.GetMember)
	g.POST("/member/create", api.CreateMember)
	g.GET("/member/list", api.GetMemberList)
	g.POST("/member/update", api.UpdateMember)
	g.POST("/member/delete", api.DeleteMember)
	// 排课
	g.POST("/course/create", api.CreateCourse)
	g.GET("/course/get", api.GetCourse)

	g.POST("/teacher/bind_course", api.BindCourse)
	g.POST("/teacher/unbind_course", api.UnBindCourse)
	g.GET("/teacher/get_course", api.GetTeacherCourse)
	g.POST("/course/schedule", api.ScheduleCourse)

	// 抢课
	g.POST("/student/book_course")
	g.GET("/student/course")
	auth := g.Group("")
	auth.Use(middleware.AuthRequired())
	{
		// 登录
		auth.POST("/auth/logout", api.Logout)
		auth.GET("/auth/whoami", api.Whoami)
>>>>>>> origin/main
	}
	return r
}
