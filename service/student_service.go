package service

import (
	"camp-course-selection/common/exception"
	"camp-course-selection/common/util"
	"camp-course-selection/model"
	"camp-course-selection/vo"
)

type StudentService struct {
}

func (m *StudentService) Course(v *vo.BookCourseRequest) util.R {
	course := model.StudentCourse{
		UserID:   v.StudentID,
		CourseID: v.CourseID,
	}
	if err := model.DB.Create(&course).Error; err != nil {
		return *util.Error(exception.UnknownError)
	}

	return *util.Ok("创建成功")
}

func (m *StudentService) GetCourse(v *vo.GetStudentCourseRequest) util.R {
	courses := []model.StudentCourse{
		UserID:   v.StudentID,
		CourseID: v.CourseID,
	}
	if err := model.DB.Where("STUDENT_ID = ?", v.StudentID).Find(&courses).Error; err != nil {
		return *util.Error(exception.UnknownError)
	}
	var courseList = make([]model.TCourse, len(courses))

	for i := 0; i < len(courses); i++ {
		if err := model.DB.Create(&courseList[i]).Error; err != nil {
			return *util.Error(exception.UnknownError)
		}
	}
	return *util.Ok(courseList)
}
