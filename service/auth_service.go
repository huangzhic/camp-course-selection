package service

import (
	"camp-course-selection/common/exception"
	"camp-course-selection/common/util"
	"camp-course-selection/model"
	"camp-course-selection/vo"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
}

// setSession 设置session
func setSession(c *gin.Context, member model.TMember) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", member.UserID)
	s.Save()
}

// Login 用户登录函数
func (m *AuthService) Login(loginVo *vo.LoginRequest, c *gin.Context) util.R {
	var member model.TMember

	model.DB.Where("USER_NAME = ?", loginVo.Username).First(&member)

	if err := model.DB.Where("USER_NAME = ?", loginVo.Username).First(&member).Error; err != nil {
		return *util.Error(exception.UserNotExisted)
	}

	if member.CheckPassword(loginVo.Password) == false {
		return *util.Error(exception.WrongPassword)
	}

	// 设置session
	setSession(c, member)

	return *util.Ok(member.UserID)
}
