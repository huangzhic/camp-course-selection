package api

import (
	"camp-course-selection/service"
	"camp-course-selection/vo"
	"github.com/gin-gonic/gin"
)

var studentService service.StudentService

//抢课
func Course(c *gin.Context) {
	var courseVo vo.BookCourseRequest
	res := vo.BookCourseResponse{}
	if err := c.ShouldBind(&courseVo); err == nil {
		res = studentService.Course(&courseVo)
		c.JSON(200, res)
	} else {
		res.Code = vo.UnknownError
		c.JSON(200, res)
	}
}

//查询
func BookCourse(c *gin.Context) {
	var courseVo vo.GetStudentCourseRequest
	res := vo.GetStudentCourseResponse{}
	if err := c.ShouldBind(&courseVo); err == nil {
		res = studentService.GetCourse(&courseVo)
		c.JSON(200, res)
	} else {
		res.Code = vo.UnknownError
		c.JSON(200, res)
	}
}
