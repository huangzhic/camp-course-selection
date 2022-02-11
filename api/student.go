package api

import (
	"camp-course-selection/common/exception"
	"camp-course-selection/common/util"
	"camp-course-selection/service"
	"camp-course-selection/vo"
	"github.com/gin-gonic/gin"
)

var studentService service.StudentService

//抢课
func Course(c *gin.Context) {
	var courseVo vo.BookCourseRequest
	if err := c.ShouldBind(&courseVo); err == nil {
		res := studentService.Course(&courseVo)
		c.JSON(200, res)
	} else {
		c.JSON(200, util.Error(exception.UnknownError))
	}
}

//查询
func BookCourse(c *gin.Context) {
	var courseVo vo.GetStudentCourseRequest
	if err := c.ShouldBind(&courseVo); err == nil {
		res := studentService.GetCourse(&courseVo)
		c.JSON(200, res)
	} else {
		c.JSON(200, util.Error(exception.UnknownError))
	}
}