package api

import (
	"camp-course-selection/common/exception"
	"camp-course-selection/common/util"
	"camp-course-selection/service"
	"camp-course-selection/vo"
	"github.com/gin-gonic/gin"
)

var courService service.CourService

//----------创建课程-------------------------------------------------
func CreateCourse(c *gin.Context) {
	var courseVo vo.CreateCourseRequest
	if err := c.ShouldBind(&courseVo); err == nil {
		res := courService.CreateCourse(&courseVo)
		c.JSON(200, res)
	} else {
		c.JSON(200, util.Error(exception.UnknownError))
	}
}

//-------------获取课程----------------------------------------------
func GetCourse(c *gin.Context) {
	var courseVo vo.GetCourseRequest
	if err := c.ShouldBind(&courseVo); err == nil {
		res := courService.GetCourse(&courseVo)
		c.JSON(200, res)
	} else {
		c.JSON(200, util.Error(exception.UnknownError))
	}
}

//-------------绑定课程---------------------------------------------
func BindCourse(c *gin.Context) {
	var bindCourse vo.BindCourseRequest
	if err := c.ShouldBind(&bindCourse); err == nil {
		res := courService.BindCourseService(&bindCourse)
		c.JSON(200, res)
	} else {
		c.JSON(200, util.Error(exception.UnknownError))
	}
}

//--------------解绑课程服务-----------------------------------------
func UnBindCourse(c *gin.Context) {
	var unBindCourse vo.UnbindCourseRequest
	if err := c.ShouldBind(&unBindCourse); err == nil {
		res := courService.UnBindCourseService(&unBindCourse)
		c.JSON(200, res)
	} else {
		c.JSON(200, util.Error(exception.UnknownError))
	}
}

//--------------获取该老师的所有课程（返回课程名称列表）------------------
func GetTeacherCourse(c *gin.Context) {
	var teacherCourse vo.GetTeacherCourseRequest
	if err := c.ShouldBind(&teacherCourse); err == nil {
		res := courService.GetTeacherCourseService(&teacherCourse)
		c.JSON(200, res)
	} else {
		c.JSON(200, util.Error(exception.UnknownError))
	}
}

//-------------排课求解器-------------------------------------------
func ScheduleCourse(c *gin.Context) {
	var schedule vo.ScheduleCourseRequest
	if err := c.ShouldBind(&schedule); err == nil {
		res := courService.ScheduleCourse(schedule)
		c.JSON(200, res)
	} else {
		c.JSON(200, util.Error(exception.UnknownError))
	}
}
