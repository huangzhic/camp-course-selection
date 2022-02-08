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

	model.DB.Where("user_name = ?", loginVo.Username).First(&member)

	if err := model.DB.Where("user_name = ?", loginVo.Username).First(&member).Error; err != nil {
		return *util.Error(exception.UserNotExisted)
	}

	if ok := member.CheckPassword(loginVo.Password); ok == false {
		//fmt.Println(loginVo.Password, "-----", member.Password)
		return *util.Error(exception.WrongPassword)
	}

	// 设置session
	setSession(c, member)
	// 设置cookie  key：camp_session    value: user_id
	//c.SetCookie("camp_session", strconv.FormatInt(int64(member.UserID), 10), 36000,
	//	"/", "localhost", false, true)
	return *util.Ok(member.UserID)
}
