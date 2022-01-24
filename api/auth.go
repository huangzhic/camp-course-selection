package api

import (
	"camp-course-selection/common/exception"
	"camp-course-selection/common/util"
	"camp-course-selection/model"
	"camp-course-selection/service"
	"camp-course-selection/vo"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var authService service.AuthService

// Login 用户登录接口

func Login(c *gin.Context) {
	var loginVo vo.LoginRequest
	if err := c.ShouldBind(&loginVo); err == nil {
		res := authService.Login(&loginVo, c)
		c.JSON(200, res)
	} else {
		c.JSON(200, util.Error(exception.UnknownError))
	}
}

// Logout 用户登出
func Logout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, &util.R{Code: 0})
}

// Whoami 获取当前用户

func Whoami(c *gin.Context) {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.TMember); ok {
			c.JSON(200, u)
		}
	}
	c.JSON(200, util.Error(exception.LoginRequired))
}
