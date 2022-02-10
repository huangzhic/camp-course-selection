package service

import (
	"camp-course-selection/common/constants"
	"camp-course-selection/common/exception"
	"camp-course-selection/common/util"
	"camp-course-selection/model"
	"camp-course-selection/vo"
	"github.com/bwmarrin/snowflake"
	"golang.org/x/crypto/bcrypt"
)

type MemberService struct {
}

// Register 用户注册

func (m *MemberService) CreateMember(memberVo *vo.CreateMemberRequest) util.R {
	member := model.TMember{
		Nickname: memberVo.Nickname,
		UserName: memberVo.Username,
		UserType: memberVo.UserType,
		Status:   constants.Active,
	}
	//检查权限
	if member.UserType != constants.Admin {
		return *util.Error(exception.PermDenied)
	}
	//检查参数是否正确
	if nick_size := len(member.Nickname); nick_size > 20 || nick_size < 4 {
		return *util.Error(exception.ParamInvalid)
	}
	if name_size := len(member.UserName); name_size > 20 || name_size < 8 {
		return *util.Error(exception.ParamInvalid)
	}
	if pass_size := len(member.Password); pass_size > 20 || pass_size < 8 {
		return *util.Error(exception.ParamInvalid)
	}

	// 表单验证
	if err := CreateMemberValid(memberVo); err != nil {
		return *err
	}
	// 雪花ID
	node, err := snowflake.NewNode(1)
	if err != nil {
		util.Log().Error("generate SnowFlakeID Error: %v", err)
		return *util.Error(exception.UnknownError)
	}
	id := node.Generate()
	member.UserID = int64(id)
	// 加密密码
	bytes, _ := bcrypt.GenerateFromPassword([]byte(memberVo.Password), bcrypt.DefaultCost)
	member.Password = string(bytes)

	// 创建用户
	if err := model.DB.Create(&member).Error; err != nil {
		return *util.Error(exception.UnknownError)
	}

	return *util.Ok(member.UserID)
}

// valid 用户注册验证表单

func CreateMemberValid(memberVo *vo.CreateMemberRequest) *util.R {
	count := int64(0)
	model.DB.Model(&model.TMember{}).Where("user_name = ?", memberVo.Username).Count(&count)
	if count > 0 {
		return util.Error(exception.UserHasExisted)
	}
	return nil
}

// 获取用户信息

func (m *MemberService) GetMember(memberVo *vo.GetMemberRequest) (model.TMember, error) {
	var member model.TMember
	result := model.DB.First(&member, memberVo.UserID)
	return member, result.Error
}

// 批量获取用户

func (m *MemberService) GetMemberList(memberVo *vo.GetMemberListRequest) ([]model.TMember, error) {
	var members []model.TMember
	result := model.DB.Find(&members)
	return members, result.Error
}

// 更新用户信息

func (m *MemberService) UpdateMember(memberVo *vo.UpdateMemberRequest) util.R {
	if err := m.MemberValid(memberVo.UserID); err != nil {
		return *err
	}
	if err := model.DB.Model(&model.TMember{}).Where("user_id = ?", memberVo.UserID).Update("nick_name", memberVo.Nickname).Error; err == nil {
		return *util.Ok(memberVo.UserID)
	} else {
		return *util.Error(exception.UnknownError)
	}
}

// 判断用户是否存在

func (m *MemberService) MemberValid(userId string) *util.R {
	count := int64(0)
	model.DB.Model(&model.TMember{}).Where("user_id=?", userId).Count(&count)
	if count == 0 {
		return util.Error(exception.UserNotExisted)
	}
	return nil
}

// 软删除

func (m *MemberService) DeleteMember(memberVo *vo.DeleteMemberRequest) util.R {
	if err := m.MemberValid(memberVo.UserID); err != nil {
		return *err
	}
	var member model.TMember
	model.DB.First(&member, memberVo.UserID)
	member.Status = 0
	if err := model.DB.Where("user_id", memberVo.UserID).Delete(&model.TMember{}).Error; err == nil {
		return *util.Ok(memberVo.UserID)
	} else {
		return *util.Error(exception.UnknownError)
	}
}
