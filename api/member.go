package api

import (
	"camp-course-selection/common/exception"
	"camp-course-selection/common/util"
	"camp-course-selection/service"
	"camp-course-selection/vo"
	"github.com/gin-gonic/gin"
)

var memberSerivce service.MemberService

// CreateMember 用户注册接口
func CreateMember(c *gin.Context) {
	var memberVo vo.CreateMemberRequest
	if err := c.ShouldBind(&memberVo); err == nil {
		res := memberSerivce.CreateMember(&memberVo)
		c.JSON(200, res)
	} else {
		c.JSON(200, util.Error(exception.UnknownError))
	}
}
