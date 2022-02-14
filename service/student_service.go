package service

import (
	"camp-course-selection/model"
	"camp-course-selection/vo"
)

type StudentService struct {
}

func (m *StudentService) BookCourse(v *vo.BookCourseRequest) (res vo.BookCourseResponse) {
	course := model.StudentCourse{
		STUDENT_ID: v.StudentID,
		COURSE_ID:  v.CourseID,
	}
	if err := model.DB.Create(&course).Error; err != nil {
		res.Code = vo.CourseHasExisted
		return
	}
	res.Code = vo.OK
	return
}

func (m *StudentService) GetStudentCourse(v *vo.GetStudentCourseRequest) (res vo.GetStudentCourseResponse) {
	courses := []model.StudentCourse{}
	if err := model.DB.Where("STUDENT_ID = ?", v.StudentID).Find(&courses).Error; err != nil {
		res.Code = vo.StudentNotExisted
		return
	}
	var courseList = make([]vo.TCourse, len(courses))
	for i := 0; i < len(courses); i++ {
		if err := model.DB.Where("course_id = ?", courses[i].COURSE_ID).Find(&courseList[i]).Error; err != nil {
			res.Code = vo.CourseNotExisted
			return
		}
	}
	res.Code = vo.OK
	res.Data.CourseList = courseList
	return
}
