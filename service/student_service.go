package service

import (
	"camp-course-selection/common/exception"
	"camp-course-selection/common/util"
	"camp-course-selection/model"
	"camp-course-selection/vo"
	"fmt"
)

type StudentService struct {
}

func (m *StudentService) Course(v *vo.BookCourseRequest) util.R {
	course := model.StudentCourse{
		STUDENT_ID: v.StudentID,
		COURSE_ID:  v.CourseID,
	}
	if err := model.DB.Create(&course).Error; err != nil {
		return *util.Error(exception.Coursed)
	}

	return *util.Ok("创建成功")
}

func (m *StudentService) GetCourse(v *vo.GetStudentCourseRequest) util.R {
	courses := []model.StudentCourse{}
	if err := model.DB.Where("STUDENT_ID = ?", v.StudentID).Find(&courses).Error; err != nil {
		return *util.Error(exception.UnknownError)
	}
	fmt.Println(courses)
	var courseList = make([]model.TCourse, len(courses))

	for i := 0; i < len(courses); i++ {
		if err := model.DB.Where("course_id = ?", courses[i].COURSE_ID).First(&courseList[i]).Error; err != nil {
			return *util.Error(exception.UnknownError)
		}
	}
	fmt.Println(courseList)
	return *util.Ok(courseList)
}
