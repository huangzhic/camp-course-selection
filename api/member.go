package api

import (
	"camp-course-selection/common/util"
	"camp-course-selection/model"
	"camp-course-selection/service"
	"camp-course-selection/vo"
	"github.com/gin-gonic/gin"
	"strconv"
)

var memberSerivce service.MemberService

// CreateMember 用户注册接口
func CreateMember(c *gin.Context) {
	res := vo.CreateMemberResponse{}
	var memberVo vo.CreateMemberRequest
	if err := c.ShouldBind(&memberVo); err == nil {
		res = memberSerivce.CreateMember(&memberVo, c)
		c.JSON(200, res)
	} else {
		res.Code = vo.UnknownError
		c.JSON(200, res)
	}
}

//GetMember 获取用户信息接口
func GetMember(c *gin.Context) {
	var memberVo vo.GetMemberRequest
	if err := c.ShouldBind(&memberVo); err == nil {
		if member, err := memberSerivce.GetMember(&memberVo); err == nil {
			if member.Status == 0 {
				c.JSON(200, util.Error(exception.UserHasDeleted))
			}
			c.JSON(200, vo.GetMemberResponse{Code: 0, Data: struct {
				UserID   string
				Nickname string
				Username string
				UserType vo.UserType
			}{UserID: strconv.FormatInt(member.UserID, 10), Nickname: member.Nickname, Username: member.UserName, UserType: vo.UserType(member.UserType)}})
		} else {
			c.JSON(200, util.Error(exception.UserNotExisted))
		}
	} else {
		c.JSON(200, util.Error(exception.UnknownError))
	}
}

//GetMemberList 批量获取成员接口
func GetMemberList(c *gin.Context) {
	var memberVo vo.GetMemberListRequest
	if err := c.ShouldBind(&memberVo); err == nil {
		if members, err := memberSerivce.GetMemberList(&memberVo); err == nil {
			c.JSON(200, vo.GetMemberListResponse{Code: 0, Data: struct{ MemberList []model.TMember }{MemberList: members}})
		} else {
			c.JSON(200, util.Error(exception.UnknownError))
		}
	}
}

//UpdateMember 更新成员数据(只允许更新昵称)接口
func UpdateMember(c *gin.Context) {
	var memberVo vo.UpdateMemberRequest
	if err := c.ShouldBind(&memberVo); err == nil {
		res := memberSerivce.UpdateMember(&memberVo)
		c.JSON(200, res)
	} else {
		c.JSON(200, util.Error(exception.UnknownError))
	}
}

//DeleteMember 删除成员（软删除）接口
func DeleteMember(c *gin.Context) {
	var memberVo vo.DeleteMemberRequest
	if err := c.ShouldBind(&memberVo); err == nil {
		res := memberSerivce.DeleteMember(&memberVo)
		c.JSON(200, res)
	} else {
		c.JSON(200, util.Error(exception.UnknownError))
	}
}
